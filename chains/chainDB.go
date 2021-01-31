// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package chains

import (
	"encoding/json"
	"fmt"
	log "github.com/emit-technology/emit-cross/log"
	"github.com/emit-technology/emit-cross/queue"
	"github.com/emit-technology/emit-cross/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"math/big"
	"path/filepath"
	"sync"
)

const (

	// minCache is the minimum amount of memory in megabytes to allocate to leveldb
	// read and write caching, split half and half.
	minCache = 16

	// minHandles is the minimum number of files handles to allocate to the open
	// database files.
	minHandles = 16
)

var (
	LAST_POLL_BLOCK_NUM = "LAST_POLL_BLOCK_NUM_%s_%d"

	LAST_VOTE_QUEUE_ID = "LAST_VOTE_QUEUE_ID_%s_%d"

	LAST_SIGN_QUEUE_ID = "LAST_SIGN_QUEUE_ID_%s_%d"

	LAST_EXECUTE_QUEUE_ID = "LAST_EXECUTE_QUEUE_ID_%s_%d"

	LAST_BATCHVOTES_QUEUE_ID = "LAST_BATCHVOTES_QUEUE_ID_%s_%d"
)

var (
	DESTINATION_PROPOSAL_MSG = "DESTINATION_PROPOSAL_MSG_%d"

	COLLECTOR_SIGN_REQ = "COLLECTOR_SIGN_REQ_%d"

	DESTINATION_EXECUTE_MSG = "DESTINATION_EXECUTE_MSG_%d"

	DESTINATION_BATCH_VOTES_MSG = "DESTINATION_BATCH_VOTES_MSG_%d"
)

type StorageSize float64

// String implements the stringer interface.
func (s StorageSize) String() string {
	if s > 1099511627776 {
		return fmt.Sprintf("%.2f TiB", s/1099511627776)
	} else if s > 1073741824 {
		return fmt.Sprintf("%.2f GiB", s/1073741824)
	} else if s > 1048576 {
		return fmt.Sprintf("%.2f MiB", s/1048576)
	} else if s > 1024 {
		return fmt.Sprintf("%.2f KiB", s/1024)
	} else {
		return fmt.Sprintf("%.2f B", s)
	}
}

type Properties struct {
	sync.RWMutex
	db *leveldb.DB
}

// Put puts the given key / value to the queue
func (p *Properties) Put(key []byte, value []byte) error {
	p.Lock()
	defer p.Unlock()
	return p.db.Put(key, value, nil)
}

func (p *Properties) Has(key []byte) (bool, error) {
	p.RLock()
	defer p.RUnlock()
	return p.db.Has(key, nil)
}

// Get returns the given key if it's present.
func (p *Properties) Get(key []byte) ([]byte, error) {
	p.RLock()
	defer p.RUnlock()
	dat, err := p.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Delete deletes the key from the queue and database
func (p *Properties) Delete(key []byte) error {
	p.Lock()
	defer p.Unlock()
	return p.db.Delete(key, nil)
}

func NewLevelDB(file string, cache int, handles int) (*leveldb.DB, error) {
	// Ensure we have some minimal caching and file guarantees
	if cache < minCache {
		cache = minCache
	}
	if handles < minHandles {
		handles = minHandles
	}
	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
		DisableSeeksCompaction: true,
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

type ChainDB struct {
	Properties       *Properties
	Msg              *queue.PrefixQueue
	CollectorChainId uint8
	log              log.Logger
}

func NewChainDB(dataDir string, cache int, handles int) (*ChainDB, error) {
	msgdb, err := NewLevelDB(filepath.Join(dataDir, "msg"), 0, 0)
	if err != nil {
		return nil, err
	}
	queue, err := queue.OpenPrefixQueue(msgdb)
	if err != nil {
		return nil, err
	}

	propertiesdb, err := NewLevelDB(filepath.Join(dataDir, "properties"), 0, 0)

	if err != nil {
		return nil, err
	}

	properties := &Properties{db: propertiesdb}

	logger := log.Root().New("db", "chanDB")
	return &ChainDB{
		Properties: properties,
		Msg:        queue,
		log:        logger,
	}, nil

}

func (c *ChainDB) UpdateNextPollBlockNum(chainId uint8, account string, blockNumber uint64) error {
	return c.Properties.Put([]byte(fmt.Sprintf(LAST_POLL_BLOCK_NUM, account, chainId)), new(big.Int).SetUint64(blockNumber).Bytes())
}

func (c *ChainDB) GetNextPollBlockNum(chainId uint8, account string) (uint64, error) {
	v, err := c.Properties.Get([]byte(fmt.Sprintf(LAST_POLL_BLOCK_NUM, account, chainId)))
	if err != nil {
		if err == errors.ErrNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return new(big.Int).SetBytes(v).Uint64(), nil
}

func (c *ChainDB) UpdateLastVoteId(chainId uint8, account string, blockNumber uint64) error {
	return c.Properties.Put([]byte(fmt.Sprintf(LAST_VOTE_QUEUE_ID, account, chainId)), new(big.Int).SetUint64(blockNumber).Bytes())
}

func (c *ChainDB) GetLastVoteId(chainId uint8, account string) (uint64, error) {
	v, err := c.Properties.Get([]byte(fmt.Sprintf(LAST_VOTE_QUEUE_ID, account, chainId)))
	if err != nil {
		if err == errors.ErrNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return new(big.Int).SetBytes(v).Uint64(), nil
}

func (c *ChainDB) UpdateLastSignId(chainId uint8, account string, blockNumber uint64) error {
	return c.Properties.Put([]byte(fmt.Sprintf(LAST_SIGN_QUEUE_ID, account, chainId)), new(big.Int).SetUint64(blockNumber).Bytes())
}

func (c *ChainDB) GetLastSignId(chainId uint8, account string) (uint64, error) {
	v, err := c.Properties.Get([]byte(fmt.Sprintf(LAST_SIGN_QUEUE_ID, account, chainId)))
	if err != nil {
		if err == errors.ErrNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return new(big.Int).SetBytes(v).Uint64(), nil
}

func (c *ChainDB) UpdateLastExecuteId(chainId uint8, account string, blockNumber uint64) error {
	return c.Properties.Put([]byte(fmt.Sprintf(LAST_EXECUTE_QUEUE_ID, account, chainId)), new(big.Int).SetUint64(blockNumber).Bytes())
}

func (c *ChainDB) GetLastExecuteId(chainId uint8, account string) (uint64, error) {
	v, err := c.Properties.Get([]byte(fmt.Sprintf(LAST_EXECUTE_QUEUE_ID, account, chainId)))
	if err != nil {
		if err == errors.ErrNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return new(big.Int).SetBytes(v).Uint64(), nil
}

func (c *ChainDB) UpdateLastBatchVotesId(chainId uint8, account string, blockNumber uint64) error {
	return c.Properties.Put([]byte(fmt.Sprintf(LAST_BATCHVOTES_QUEUE_ID, account, chainId)), new(big.Int).SetUint64(blockNumber).Bytes())
}

func (c *ChainDB) GetLastBatchVotesId(chainId uint8, account string) (uint64, error) {
	v, err := c.Properties.Get([]byte(fmt.Sprintf(LAST_BATCHVOTES_QUEUE_ID, account, chainId)))
	if err != nil {
		if err == errors.ErrNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return new(big.Int).SetBytes(v).Uint64(), nil
}

func (c *ChainDB) AddProposalMsg(transfer types.FTTransfer) error {
	_, err := c.Msg.EnqueueObjectAsJSON([]byte(fmt.Sprintf(DESTINATION_PROPOSAL_MSG, transfer.DestinationId)), transfer)
	return err
}

func (c *ChainDB) GetProposalMsgById(chainId uint8, id uint64) (*types.FTTransfer, error) {
	item, err := c.Msg.PeekByID([]byte(fmt.Sprintf(DESTINATION_PROPOSAL_MSG, chainId)), id)
	if err != nil {
		return nil, err
	}
	var result types.FTTransfer
	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func (c *ChainDB) AddSignReq(transfer types.FTTransfer) error {
	_, err := c.Msg.EnqueueObjectAsJSON([]byte(fmt.Sprintf(COLLECTOR_SIGN_REQ, c.CollectorChainId)), transfer)
	return err
}

func (c *ChainDB) GetSignReqById(id uint64) (*types.FTTransfer, error) {
	item, err := c.Msg.PeekByID([]byte(fmt.Sprintf(COLLECTOR_SIGN_REQ, c.CollectorChainId)), id)
	if err != nil {
		return nil, err
	}
	var result types.FTTransfer
	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func (c *ChainDB) AddExecuteMsg(transfer types.FTTransfer) error {
	_, err := c.Msg.EnqueueObjectAsJSON([]byte(fmt.Sprintf(DESTINATION_EXECUTE_MSG, transfer.DestinationId)), transfer)
	return err
}

func (c *ChainDB) GetExecuteMsgById(chainId uint8, id uint64) (*types.FTTransfer, error) {
	item, err := c.Msg.PeekByID([]byte(fmt.Sprintf(DESTINATION_EXECUTE_MSG, chainId)), id)
	if err != nil {
		return nil, err
	}
	var result types.FTTransfer
	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func (c *ChainDB) AddBatchVotesMsg(votes types.BatchVotes) error {
	_, err := c.Msg.EnqueueObjectAsJSON([]byte(fmt.Sprintf(DESTINATION_BATCH_VOTES_MSG, votes.DestinationId)), votes)
	return err
}

func (c *ChainDB) GetBatchVotesById(chainId uint8, id uint64) (*types.BatchVotes, error) {
	item, err := c.Msg.PeekByID([]byte(fmt.Sprintf(DESTINATION_BATCH_VOTES_MSG, chainId)), id)
	if err != nil {
		return nil, err
	}
	var result types.BatchVotes
	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
