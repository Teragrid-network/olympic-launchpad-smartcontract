package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"time"
)

func main() {
	loadEnvErr := godotenv.Load()
	if loadEnvErr != nil {
		logrus.Error(loadEnvErr)
	}

	client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	handleErr(err)
	deploy(client)
}

func deploy(client *ethclient.Client) {
	tree := setupMerkleTree()
	root := tree.Root()
	logrus.Info("Root primitives: ", root)
	logrus.Info("Root: ", common.BytesToHash(root).Hex())
	bData := common.HexToAddress(os.Getenv("WHITELIST_ACCOUNT1")).Bytes()
	proofGet, err := tree.GenerateProof(bData)
	handleErr(err)
	logrus.Info("Proof: ", proofGet)
	root32 := convertByteToByte32(root)
	var proof32 = make([][32]byte, 0)
	for _, v := range proofGet.Hashes {
		converted := convertByteToByte32(v)
		proof32 = append(proof32, converted)
	}

	logrus.Info("root after comverted: ", convertByteToByte32(root))

	chainId, retrieveChainIDErr := client.ChainID(context.Background())
	handleErr(retrieveChainIDErr)
	users := accountBinding(chainId)
	auth := users[0]
	var tmpProof = StakingProof{
		Hashes: proof32,
		Index:  big.NewInt(int64(proofGet.Index)),
	}
	logrus.Info("Staking proof: ", tmpProof)
	isVerify, _ := VerifyProofUsing(bData, proofGet, root)
	logrus.Info("verify: ", isVerify)
	addressMUUV, _, _, err := DeployMUUV(auth, client)
	handleErr(err)
	logrus.Infof("MUUV address: %s", addressMUUV.Hex())
	token, err := NewMUUV(addressMUUV, client)
	handleErr(err)
	enableStakeDate := time.Now().Unix() + 30
	disableStakingDate := time.Now().Unix() + 1e10
	addressStaking, trans, _, err := DeployStaking(auth, client, addressMUUV, big.NewInt(enableStakeDate), big.NewInt(disableStakingDate), root32) // add param merkle Root
	handleErr(err)
	logrus.Infof("Staking address: %s", addressStaking.Hex())
	logrus.Infof("Staking transaction: %s", trans.Hash())
	time.Sleep(time.Second * 5)

	staking(token, addressStaking, client, auth, tmpProof)
}
func staking(token *MUUV, addressStaking common.Address, client *ethclient.Client, auth *bind.TransactOpts, proof StakingProof) {
	staking, err := NewStaking(addressStaking, client)
	merkleroot, err := staking.MerkleRoot(nil)
	handleErr(err)
	logrus.Info("MerkleRoot in contract: ", merkleroot)
	time.Sleep(time.Second * 5)
	balance, err := token.BalanceOf(nil, auth.From)
	logrus.Info("Call from: ", auth.From)
	handleErr(err)
	logrus.Info("BalanceOf: ", balance)

	//Transfer from owner to the contract
	logrus.Info("Transfer from owner to contract....")
	_, err = token.Transfer(auth, addressStaking, big.NewInt(1000000)) // 1 mil
	handleErr(err)
	time.Sleep(time.Second * 5)

	// Approve
	logrus.Info("Approving....")
	logrus.Info("address Staking: ", addressStaking)
	_, err = token.Approve(auth, addressStaking, big.NewInt(100000)) // 1 mil
	handleErr(err)
	time.Sleep(time.Second * 20)

	//Staking
	stakingTx, err := staking.Stake(auth, big.NewInt(10000), big.NewInt(50), big.NewInt(2), big.NewInt(1))
	handleErr(err)
	logrus.Info("Staking: ", stakingTx.Hash())
	time.Sleep(time.Second * 5)

	//Get Staking Info
	stakingInfo, err := staking.StakeInfo(nil, auth.From, big.NewInt(0))
	handleErr(err)
	logrus.Info("Amount: ", stakingInfo.Amount.String())
	logrus.Info("Amount Claimed: ", stakingInfo.AmountClaimed.String())
	logrus.Info("Last Milestones: ", stakingInfo.Lastmilestones.String())
	logrus.Info("Milestone: ", stakingInfo.Milestones.String())
	logrus.Info("Total Reward: ", stakingInfo.TotalReward.String())

	claim(staking, auth, stakingTx, proof)
}

// @param: auth: caller to contract
func claim(staking *Staking, auth *bind.TransactOpts, stakingTx *types.Transaction, merkleProof StakingProof) {
	//Claim
	claimTx, err := staking.Claim(auth, big.NewInt(0), merkleProof)
	handleErr(err)
	logrus.Info("Claim: ", claimTx.Hash())
	time.Sleep(time.Second * 5)
	logrus.Info("call from: ", auth.From)
	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(30), big.NewInt(2), big.NewInt(1))
	handleErr(err)
	time.Sleep(time.Second * 5)
	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(20), big.NewInt(9), big.NewInt(1))
	handleErr(err)
	time.Sleep(time.Second * 5)
	stakingTx, err = staking.Stake(auth, big.NewInt(10000), big.NewInt(10), big.NewInt(4), big.NewInt(1))
	handleErr(err)

	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{
	//		addressMUUV,
	//		addressStaking,
	//	},
	//}
	//go eventListener(client, query)
}

func accountBinding(chainID *big.Int) []*bind.TransactOpts {
	privateKeyMain, err := crypto.HexToECDSA(os.Getenv("PRIVATEKEY1"))
	handleErr(err)
	privateKey1, err := crypto.HexToECDSA(os.Getenv("PRIVATEKEY2"))
	handleErr(err)
	privateKey2, err := crypto.HexToECDSA(os.Getenv("PRIVATEKEY3"))
	handleErr(err)
	user1, err := bind.NewKeyedTransactorWithChainID(privateKeyMain, chainID)
	user2, err := bind.NewKeyedTransactorWithChainID(privateKey1, chainID)
	user3, err := bind.NewKeyedTransactorWithChainID(privateKey2, chainID)
	return []*bind.TransactOpts{user1, user2, user3}

}
func handleErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func setupMerkleTree() *MerkleTree {
	// Set up whitelist
	var whitelist = [][]byte{
		common.HexToAddress(os.Getenv("WHITELIST_ACCOUNT1")).Bytes(),
		common.HexToAddress(os.Getenv("WHITELIST_ACCOUNT2")).Bytes(),
		common.HexToAddress(os.Getenv("WHITELIST_ACCOUNT3")).Bytes(),
		common.HexToAddress(os.Getenv("WHITELIST_ACCOUNT5")).Bytes(),
	}
	tree, _ := NewUsing(whitelist)
	logrus.Info("address in crypto hash keccak: ", crypto.Keccak256Hash(whitelist[3]))
	//logrus.Info("Proof1: ", proof)
	//logrus.Info("Merkle Proof2: ", hex.EncodeToString(proof2.Hashes[0]))
	//logrus.Info("Proof2: ", proof2)

	return tree
}
func convertByteToByte32(data []byte) [32]byte {
	var data32 [32]byte
	if len(data) == 0 {
		return data32
	}
	for i := range data32 {
		data32[i] = data[i]
	}
	return data32
	//return *(*[32]byte)(unsafe.Pointer(&data))
}

//0x5931b4ed56ace4c46b68524cb5bcbf4195f1bbaacbe5228fbd090546c88dd229
func cloneKecaak256ABIEncode() {
	hash := crypto.Keccak256Hash(common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4").Bytes())
	logrus.Info("hash: ", hash.Hex())
}
