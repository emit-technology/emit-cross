// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the cross application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package main

import (
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/chains/collector"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/emit-technology/emit-cross/chains/ethereum"

	"github.com/emit-technology/emit-cross/chains/sero"

	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/metrics/health"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	"github.com/emit-technology/emit-cross/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	config.ConfigFileFlag,
	config.VerbosityFlag,
	config.KeystorePathFlag,
	config.KeyFileFlag,
	config.BlockstorePathFlag,
	config.FreshStartFlag,
	config.LatestBlockFlag,
	config.MetricsFlag,
	config.MetricsPort,
}

var generateFlags = []cli.Flag{
	config.PasswordFlag,
}

var devFlags = []cli.Flag{
	config.TestKeyFlag,
}

var importFlags = []cli.Flag{
	config.EthereumImportFlag,
	config.PrivateKeyFlag,
	config.PasswordFlag,
}

var inspectFlags = []cli.Flag{
	config.KeyFileFlag,
	config.PasswordFlag,
}

var accountCommand = cli.Command{
	Name:  "accounts",
	Usage: "manage bridge keystore",
	Description: "The accounts command is used to manage the bridge keystore.\n" +
		"\tTo generate a new account (key type generated is determined on the flag passed in): cross accounts generate\n" +
		"\tTo import a keystore file: cross accounts import path/to/file\n" +
		"\tTo import a geth keystore file: cross accounts import --ethereum path/to/file\n" +
		"\tTo import a private key file: cross accounts import --privateKey private_key\n" +
		"\tTo list keys: cross accounts list",
	Subcommands: []*cli.Command{
		{
			Action: wrapHandler(handleGenerateCmd),
			Name:   "generate",
			Usage:  "generate bridge keystore, key type determined by flag",
			Flags:  generateFlags,
			Description: "The generate subcommand is used to generate the bridge keystore.\n" +
				"\tIf no options are specified, a secp256k1 key will be made.",
		},
		{
			Action: wrapHandler(handleImportCmd),
			Name:   "import",
			Usage:  "import bridge keystore",
			Flags:  importFlags,
			Description: "The import subcommand is used to import a keystore for the bridge.\n" +
				"\tA path to the keystore must be provided\n" +
				"\tUse --ethereum to import an ethereum keystore from external sources such as geth\n" +
				"\tUse --privateKey to create a keystore from a provided private key.",
		},
		{
			Action:      wrapHandler(handleListCmd),
			Name:        "list",
			Usage:       "list bridge keystore",
			Description: "The list subcommand is used to list all of the bridge keystores.\n",
		},
		{
			Action:      wrapHandler(handleInspectCmd),
			Name:        "inspect",
			Usage:       "inspect a keystore",
			Flags:       inspectFlags,
			Description: "The inspect subcommand is used to inspect the bridge keystores.\n",
		},
	},
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2020 EMIT Foundation"
	app.Name = "emit-cross"
	app.Usage = "emit-cross"
	app.Authors = []*cli.Author{{Name: "EMIT Foundation 2020"}}
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&accountCommand,
	}

	app.Flags = append(app.Flags, cliFlags...)
	app.Flags = append(app.Flags, devFlags...)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func startLogger(ctx *cli.Context) error {
	logger := log.Root()
	handler := logger.GetHandler()
	var lvl log.Lvl

	if lvlToInt, err := strconv.Atoi(ctx.String(config.VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.String(config.VerbosityFlag.Name)); err != nil {
		return err
	}
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))

	return nil
}

func getDefaultPath(pathPostfix string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, pathPostfix), nil
}

func run(ctx *cli.Context) error {
	err := startLogger(ctx)
	if err != nil {
		return err
	}

	log.Info("Starting ChainBridge...")

	cfg, err := config.GetConfig(ctx)
	if err != nil {
		return err
	}

	// Check for test key flag
	var ks string
	var kf string
	var insecure bool
	if key := ctx.String(config.TestKeyFlag.Name); key != "" {
		ks = key
		insecure = true
	} else {
		ks = cfg.KeystorePath
		kf = cfg.KeyFile
		if kf == "" {
			return errors.New("the file name of the kestore must be specified  with --keyFile")
		}
	}

	// Used to signal core shutdown due to fatal error
	sysErr := make(chan error)
	c := core.NewCore(sysErr)

	blockStorePath := ctx.String(config.BlockstorePathFlag.Name)
	if blockStorePath == "" {
		blockStorePath, _ = getDefaultPath(".emit-cross/blockstore")
	}

	if len(cfg.Chains) < 3 {
		return errors.New("invalid chains config")
	}
	var collectorChain *collector.Chain
	var crossChains []config.RawChainConfig
	for _, chain := range cfg.Chains {
		if chain.Type == "collector" {

			chainConfig := &core.ChainConfig{
				Name:           chain.Name,
				Id:             msg.ChainId(chains.CollectorChainId),
				Endpoint:       chain.Endpoint,
				From:           kf,
				KeystorePath:   ks,
				Insecure:       insecure,
				BlockstorePath: blockStorePath,
				FreshStart:     ctx.Bool(config.FreshStartFlag.Name),
				LatestBlock:    ctx.Bool(config.LatestBlockFlag.Name),
				Opts:           chain.Opts,
			}

			var m *metrics.ChainMetrics

			logger := log.Root().New("chain", chainConfig.Name)

			if ctx.Bool(config.MetricsFlag.Name) {
				m = metrics.NewChainMetrics(chain.Name)
			}

			collectorChain, err = collector.InitializeChain(chainConfig, logger, sysErr, m)

			if err != nil {
				return err
			}
			c.AddChain(collectorChain)

		} else {
			crossChains = append(crossChains, chain)
		}
	}
	if collectorChain == nil {
		return errors.New("not config collector")
	}

	for _, chain := range crossChains {
		chainId, errr := strconv.Atoi(chain.Id)
		if errr != nil {
			return errr
		}
		if chainId == 0 {
			return errors.New("only collector chainId is 0")
		}
		chainConfig := &core.ChainConfig{
			Name:           chain.Name,
			Id:             msg.ChainId(chainId),
			Endpoint:       chain.Endpoint,
			From:           kf,
			KeystorePath:   ks,
			Insecure:       insecure,
			BlockstorePath: blockStorePath,
			FreshStart:     ctx.Bool(config.FreshStartFlag.Name),
			LatestBlock:    ctx.Bool(config.LatestBlockFlag.Name),
			Opts:           chain.Opts,
		}
		var newChain core.Chain
		var propState chains.ProposalState
		var m *metrics.ChainMetrics

		logger := log.Root().New("chain", chainConfig.Name)

		if ctx.Bool(config.MetricsFlag.Name) {
			m = metrics.NewChainMetrics(chain.Name)
		}

		if chain.Type == "ethereum" {
			newChain, propState, err = ethereum.InitializeChain(chainConfig, logger, sysErr, m)
		} else if chain.Type == "sero" {
			newChain, propState, err = sero.InitializeChain(chainConfig, logger, sysErr, m)
		} else {
			return errors.New("unrecognized Chain Type")
		}

		if err != nil {
			return err
		}
		c.AddChain(newChain)
		collectorChain.AddPropState(uint8(chainId), propState)

	}

	// Start prometheus and health server
	if ctx.Bool(config.MetricsFlag.Name) {
		port := ctx.Int(config.MetricsPort.Name)
		blockTimeoutStr := os.Getenv(config.HealthBlockTimeout)
		blockTimeout := config.DefaultBlockTimeout
		if blockTimeoutStr != "" {
			blockTimeout, err = strconv.ParseInt(blockTimeoutStr, 10, 0)
			if err != nil {
				return err
			}
		}
		h := health.NewHealthServer(port, c.Registry, int(blockTimeout))

		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.HandleFunc("/health", h.HealthStatus)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("Health status server is shutting down", err)
			} else {
				log.Error("Error serving metrics", "err", err)
			}
		}()
	}

	c.Start()

	return nil
}
