package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"olympic-launchpad/api"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load(".env") //TODO

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestScanBlockchain(t *testing.T) {

	client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	defer client.Close()
	defer fmt.Println("close")
	if err != nil {
		log.Fatal(err)
	}

	scanTransByContractAddress(common.HexToHash("0x1b7a03d2b166e26bbaaa170fc2e22430fb93b32515b3f0bc13ada4c5895dc3d9"), client)

}

//TODO: solc, abigen terminal readme; openzeppelin node module install
func TestDeployReadWriteContract(t *testing.T) { //TODO: bad handshake error

	//TODO: Golang time traveller
	//TODO: AutoRemapping
	//TODO: Deployed check
	//TODO: staking.StakingFilterer
	//TODO: staking.ParseStaked()

	client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc") // use websocket if scan events: wss://api.avax-test.network/ext/bc/C/ws
	printErr(err)

	chainID, err := client.ChainID(context.Background())
	printErr(err)

	user := accountHandle(chainID)
	auth := user[0]

	// Deploy TokenMUUV
	addressMUUV, _, _, err := api.DeployMUUV(auth, client)
	printErr(err)
	fmt.Println("Deploy MUUV at:", addressMUUV.Hex())
	token, err := api.NewMUUV(addressMUUV, client)
	printErr(err)
	time.Sleep(time.Second * 2)

	//Deploy Staking
	enableStakeDate := time.Now().Unix() + 7             // right now + x second
	disableStakeDate := time.Now().Unix() + 100000000000 // infinity
	addressStaking, trans, _, err := api.DeployStaking(auth, client, addressMUUV, big.NewInt(enableStakeDate), big.NewInt(disableStakeDate))
	printErr(err)
	fmt.Println("Deploy Staking at:", addressStaking.Hex())
	fmt.Println("Deploy Staking Hash:", trans.Hash())
	staking, err := api.NewStaking(addressStaking, client)
	printErr(err)
	time.Sleep(time.Second * 2)

	//EVENT
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{ /*addressMUUV,*/ addressStaking},
	//}
	//go eventListener(client, query)

	//Get MUUV token balance
	balance, err := token.BalanceOf(nil, auth.From)
	printErr(err)
	fmt.Println("BalanceOf: ", balance)

	//Transfer from owner to the contract
	fmt.Println("Transfer from owner to contract....")
	_, err = token.Transfer(auth, addressStaking, big.NewInt(1000000)) // 1 mil
	printErr(err)
	time.Sleep(time.Second * 5)

	// Approve
	fmt.Println("Approving....")
	_, err = token.Approve(auth, addressStaking, big.NewInt(100000)) // 1 mil
	printErr(err)
	time.Sleep(time.Second * 5)

	//Staking
	stakingTx, err := staking.Stake(auth, big.NewInt(10000), big.NewInt(50), big.NewInt(2), big.NewInt(1))
	printErr(err)
	fmt.Println("Staking: ", stakingTx.Hash())
	time.Sleep(time.Second * 5)

	//Get Staking Info
	stakingInfo, err := staking.StakeInfo(nil, auth.From, big.NewInt(0))
	printErr(err)
	fmt.Println("Amount: ", stakingInfo.Amount.String())
	fmt.Println("Amount Claimed: ", stakingInfo.AmountClaimed.String())
	fmt.Println("Last Milestones: ", stakingInfo.Lastmilestones.String())
	fmt.Println("Milestone: ", stakingInfo.Milestones.String())
	fmt.Println("Total Reward: ", stakingInfo.TotalReward.String())

	//Claim
	claimtx, err := staking.Claim(auth, big.NewInt(0))
	printErr(err)
	fmt.Println("Claim: ", claimtx.Hash())
	time.Sleep(time.Second * 5)

	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(30), big.NewInt(2), big.NewInt(1))
	printErr(err)
	time.Sleep(time.Second * 5)
	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(20), big.NewInt(9), big.NewInt(1))
	printErr(err)
	time.Sleep(time.Second * 5)
	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(10), big.NewInt(4), big.NewInt(1))
	printErr(err)
}

func getSignatureHash(functionName string) []byte {
	var contractAbi abi.ABI

	contractAbi, _ = abi.JSON(strings.NewReader(api.StakingMetaData.ABI))

	inputs := contractAbi.Methods[functionName].Inputs
	funcArgsType := make([]string, 0) // [] => [""] => ["", ""] // re-allocate
	for _, input := range inputs {
		funcArgsType = append(funcArgsType, input.Type.String())
	}

	functionSignature := fmt.Sprintf("%v(%v)", // output: FunctionName(argsType1,argsTyp2...)
		contractAbi.Methods[functionName].Name,
		strings.Join(funcArgsType, ","),
	)
	return crypto.Keccak256([]byte(functionSignature))[:4] ///0x1234
}

func eventListener(client *ethclient.Client, query ethereum.FilterQuery) {

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	printErr(err)

	contractAbi, err := abi.JSON(strings.NewReader(api.StakingMetaData.ABI))
	printErr(err)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if len(vLog.Data) == 0 {
				fmt.Println("LOG WITH NO DATA!!!!")
				continue
			}
			fmt.Println("-----------------LOG")
			fmt.Println("Address: ", vLog.Address) // pointer to event log
			fmt.Println("BlockNumber: ", vLog.BlockNumber)
			fmt.Println("Topics: ", vLog.Topics)
			fmt.Println("Txhash: ", vLog.TxHash)

			data, err := contractAbi.Unpack("Staked", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Data from event: ", data)
			fmt.Println("------------------------------------------------")
		}
	}
}

func printErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//
func accountHandle(chainId *big.Int) []*bind.TransactOpts {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	pk1, err := crypto.HexToECDSA(os.Getenv("PRIKEY1"))
	printErr(err)
	pk2, err := crypto.HexToECDSA(os.Getenv("PRIKEY2"))
	printErr(err)
	pk3, err := crypto.HexToECDSA(os.Getenv("PRIKEY3"))
	acc1, _ := bind.NewKeyedTransactorWithChainID(pk1, chainId)
	acc2, _ := bind.NewKeyedTransactorWithChainID(pk2, chainId)
	acc3, _ := bind.NewKeyedTransactorWithChainID(pk3, chainId)

	return []*bind.TransactOpts{acc1, acc2, acc3}
}

func scanBlock(_currentScannedBlock int64, contractAddress common.Address, client *ethclient.Client) [][]interface{} {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(_currentScannedBlock))
	printErr(err)
	fmt.Println("Block:", _currentScannedBlock)
	// check if scanned contracts is equal to given contract
	var result [][]interface{}
	for _, tx := range block.Transactions() {
		if tx.To() != nil {
			if *tx.To() == contractAddress {
				contractAbi, err := abi.JSON(strings.NewReader(string(api.StakingMetaData.ABI)))
				printErr(err)

				if method, ok := contractAbi.Methods["stake"]; ok {
					if string(tx.Data()[:4]) == string(getSignatureHash("stake")) { //methodId
						if params, err := method.Inputs.Unpack(tx.Data()[4:]); err == nil {
							//Get data here
							fmt.Println("Stake: ", params)
							result = append(result, params)
						}
					}
				}
				if method, ok := contractAbi.Methods["claim"]; ok {
					if string(tx.Data()[:4]) == string(getSignatureHash("claim")) { //methodId
						if params, err := method.Inputs.Unpack(tx.Data()[4:]); err == nil {
							fmt.Println("Claim: ", params)
							result = append(result, params)
						}
					}
				}

			} else {
				//fmt.Println("doesn't found trans in this block")
			}
		}
	}
	fmt.Println(result)
	return result
}

func scanTransByContractAddress(txHash common.Hash, client *ethclient.Client) {

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	printErr(err)

	contractAddress := receipt.ContractAddress

	currentScannedBlock, _ := strconv.ParseInt(receipt.BlockNumber.String(), 10, 64)

	//TICKER

	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			for {
				currentBlock, _ := client.BlockNumber(context.Background())
				if int64(currentBlock) <= currentScannedBlock {
					break
				}
				go scanBlock(currentScannedBlock, contractAddress, client)
				currentScannedBlock++
			}
		}
	}

	fmt.Println(client.BlockNumber(context.Background()))
}

func getTransByBlockNumber(_blockNumber int64, client *ethclient.Client) int {

	blockNumber := big.NewInt(_blockNumber)

	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	//for _, tx := range block.Transactions() {
	//	fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	//	fmt.Println("To: ", tx.To())
	//	fmt.Printf("TX Value: %s\n", tx.Value().String())
	//	fmt.Printf("TX Gas: %d\n", tx.Gas())
	//	fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
	//	fmt.Printf("TX Nonce: %d\n", tx.Nonce())
	//	fmt.Printf("TX Data: %v\n", tx.Data())
	//
	//}
	return len(block.Transactions())
}

func getTranDetail(tranHash string, client *ethclient.Client) *common.Address {
	// Grab a transaction by it's individual hash
	txHash := common.HexToHash(tranHash)
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("TX Value: %s\n", tx.Value().String())
	//fmt.Printf("TX Gas: %d\n", tx.Gas())
	//fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
	//fmt.Printf("TX Nonce: %d\n", tx.Nonce())
	//fmt.Printf("TX Data: %v\n", tx.Data())
	//fmt.Printf("TX To: %s\n", tx.To().String())
	//fmt.Printf("Pending?: %v\n", isPending)
	return tx.To()
}
