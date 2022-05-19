package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_ConnectClient(t *testing.T) {
	client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	assert.Equal(t, nil, err)

	chainId, err := client.ChainID(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, "43113", chainId.String())
}

func Test_GetTransaction(t *testing.T) {
	tx, isPending, err := client.TransactionByHash(context.Background(),
		common.HexToHash("0x78c0d4c02ae4523e0ce565d959fc432213e0bfcf11576b5ca169de018e890ef1"))
	assert.Nil(t, err)
	assert.False(t, isPending)
	assert.NotNil(t, tx)
	assert.Equal(t, types.LegacyTxType, int(tx.Type()))
	assert.NotNil(t, tx.To()) // to == nil => deploy smart contract

	wei := tx.Value()
	avax := wei.Div(wei, big.NewInt(params.Ether))
	assert.Equal(t, int64(10), avax.Int64())
}

func Test_GetLogs(t *testing.T) {
	blockHash := common.HexToHash("0x8d951b09cda20bcf015ef1191915cdcf2c9cd6d6e076a918c981b92ee7257a7b")
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: []common.Address{common.HexToAddress("0xfc7c98ff48aa50d75b77a3ca3e7f528817b88255")},
	})
	assert.Equal(t, nil, err)
	assert.NotNil(t, logs)
}
