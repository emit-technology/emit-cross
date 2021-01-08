# Emit-Cross
  
  ###### This is the core software that the relays running between block-chains, which supporting tokens to cross between ETH and SERO chains. 

    
# Installation

## Dependencies

- [libsuperzk](https://github.com/sero-cash/go-czero-import): 
Used for sero account.You must configure the environment variables before running

    ```
   git clone https://github.com/sero-cash/go-czero-import
  
  cd go-czero-import 
  
  linux:
    
        export LD_LIBRARY_PATH="./czero/lib_LINUX_AMD64_V3"
  
  darwin:
  
        export DYLD_LIBRARY_PATH="./czero/lib_DARWIN_AMD64"

    ```
   
   If ubuntu, you may also need to install libgmp
   
   ```
        apt-get install libgmp-dev    or
        
        apt-get install libgomp1
  ```


## Building

`make build`: Builds `cross` in `./build`.

**or**

`make install`: Uses `go install` to add `cross` to your GOBIN.

# Configuration

> Note: TOML configs have been deprecated in favour of JSON

A chain configurations take this form:

```
{
    "name": "eth",                      // Human-readable name
    "type": "ethereum",                 // Chain type (eg. "ethereum" or "substrate")
    "id": "0",                          // Chain ID
    "endpoint": "ws://<host>:<port>",   // Node endpoint
    "from": "0xff93...",                // On-chain address of relayer
    "opts": {},                         // Chain-specific configuration options (see below)
}
```

See `config.json.example` for an example configuration. 

### Ethereum Options

Ethereum chains support the following additional options:

```
{
    "bridge": "0x12345...",          // Address of the bridge contract (required)
    "erc20Handler": "0x1234...",     // Address of erc20 handler (required)
    "maxGasPrice": "0x1234",         // Gas price for transactions (default: 20000000000)
    "gasLimit": "0x1234",            // Gas limit for transactions (default: 6721975)
    "http": "true",                  // Whether the chain connection is ws or http (default: false)
    "startBlock": "1234",            // The block to start processing events from (default: 0)
    "blockConfirmations": "10"       // Number of blocks to wait before processing a block
    "commitNode":"true"              //  Whether to submit an action to execute the proposal

}
```

### Sero Options

sero supports the following additonal options:

```
{
     "accountEndpoint":"http://<host>:<port", //account endpoint 
}
```

## Blockstore

The blockstore is used to record the last block the relayer processed, so it can pick up where it left off. 

If a `startBlock` option is provided (see [Configuration](#configuration)), then the greater of `startBlock` and the latest block in the blockstore is used at startup.

To disable loading from the blockstore specify the `--fresh` flag. A custom path for the blockstore can be provided with `--blockstore <path>`. For development, the `--latest` flag can be used to start from the current block and override any other configuration.

## Keystore

ChainBridge requires keys to sign and submit transactions, and to identify each bridge node on chain.

To use secure keys, see `ross accounts --help`. The keystore password can be supplied with the `KEYSTORE_PASSWORD` environment variable.

To import external ethereum keys, such as those generated with geth, use `cross accounts import --ethereum /path/to/key`.

To import private keys as keystores, use `cross accounts import --privateKey key`.

To inspect keystores use `cross accounts --keystore path/to/keystore inspect --keyFile filename`

For testing purposes, cross provides 5 test keys. The can be used with `--testkey <name>`, where `name` is one of `Alice`, `Bob`, `Charlie`, `Dave`, or `Eve`. 
