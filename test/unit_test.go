package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	_ "testing/quick"
)

var client, _ = ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")

func Test_GetSignatureHash(t *testing.T) {
	a := getSignatureHash("stake")
	assert.Equal(t, byte(206), a[0])
	assert.Equal(t, byte(147), a[1])
	assert.Equal(t, byte(168), a[2])
	assert.Equal(t, byte(130), a[3])
	assert.Equal(t, 4, len(a))
	a = getSignatureHash("claim")
	assert.Equal(t, byte(55), a[0])
	assert.Equal(t, byte(150), a[1])
	assert.Equal(t, byte(7), a[2])
	assert.Equal(t, byte(245), a[3])
	assert.Equal(t, 4, len(a))
}
func Test_EventListener(t *testing.T) { //TODO

	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{ /*addressMUUV,*/ /*addressStaking*/},
	//}
	//eventListener(client, query)
}
func Test_AccountHandle(t *testing.T) {
	chainID, err := client.ChainID(context.Background())
	printErr(err)
	accounts := accountHandle(chainID)
	fmt.Println(len(accounts))
	assert.Equal(t, len(accounts), 3)
	assert.NotNil(t, accounts[0])
	assert.NotNil(t, accounts[1])
	assert.NotNil(t, accounts[2])

}
func Test_ScanBlock(t *testing.T) {
	a := scanBlock(9049017, common.HexToAddress("0x8d40ec7eb57b3edd766d66ff108e29668e208956"), client)
	assert.Equal(t, len(a[0]), 4)
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0][0], big.NewInt(10000))
	assert.Equal(t, a[0][1], big.NewInt(10))
	assert.Equal(t, a[0][2], big.NewInt(4))
	assert.Equal(t, a[0][3], big.NewInt(1))
}
func Test_ScanTransByContractAddress(t *testing.T) { //TODO

}
func Test_GetTransByBlockNumber(t *testing.T) {
	transCount := getTransByBlockNumber(8940888, client)
	assert.Equal(t, transCount, 3)
	transCount = getTransByBlockNumber(8940889, client)
	assert.Equal(t, transCount, 2)
	transCount = getTransByBlockNumber(8940890, client)
	assert.Equal(t, transCount, 1)
	transCount = getTransByBlockNumber(8940891, client)
	assert.Equal(t, transCount, 1)
	transCount = getTransByBlockNumber(8940892, client)
	assert.Equal(t, transCount, 2)

}
func Test_GetTranDetail(t *testing.T) {
	to := getTranDetail("0x35ca2c524c0e2d5195c364ee699d803902afe5ad17e620f65e4f45d832ff5090", client)
	assert.Equal(t, *to, common.HexToAddress("0xe50256721f837e6773befdc2a072eefedae99e27"))
}
