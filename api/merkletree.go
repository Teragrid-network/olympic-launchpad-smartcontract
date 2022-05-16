package main

import (
	"bytes"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
)

type MerkleTree struct {
	// salt is the optional salt hashed with data to avoid rainbow attacks
	// data is the data from which the Merkle tree is created
	data [][]byte
	// nodes are the leaf and branch nodes of the Merkle tree
	nodes [][]byte
}
type Proof struct {
	Hashes [][]byte
	Index  uint64
}

func newProof(hashes [][]byte, index uint64) *Proof {
	return &Proof{
		Hashes: hashes,
		Index:  index,
	}
}

func NewUsing(data [][]byte) (*MerkleTree, error) {
	if len(data) == 0 {
		return nil, errors.New("tree must have at least 1 piece of data")
	}

	branchesLen := int(math.Exp2(math.Ceil(math.Log2(float64(len(data))))))

	// We pad our data length up to the power of 2
	nodes := make([][]byte, branchesLen+len(data)+(branchesLen-len(data)))
	// Leaves
	for i := range data {
		nodes[i+branchesLen] = crypto.Keccak256Hash(data[i]).Bytes()

	}
	// Branches
	for i := branchesLen - 1; i > 0; i-- {
		nodes[i] = crypto.Keccak256Hash(append(nodes[i*2], nodes[i*2+1]...)).Bytes()
	}

	tree := &MerkleTree{
		nodes: nodes,
		data:  data,
	}

	return tree, nil
}
func (t *MerkleTree) GenerateProof(data []byte) (*Proof, error) {
	// Find the index of the data
	index, err := t.indexOf(data)
	if err != nil {
		return nil, err
	}

	proofLen := int(math.Ceil(math.Log2(float64(len(t.data)))))
	hashes := make([][]byte, proofLen)

	cur := 0
	for i := index + uint64(len(t.nodes)/2); i > 1; i /= 2 {
		hashes[cur] = t.nodes[i^1]
		cur++
	}
	return newProof(hashes, index), nil
}
func (t *MerkleTree) indexOf(input []byte) (uint64, error) {
	for i, data := range t.data {
		if bytes.Compare(data, input) == 0 {
			return uint64(i), nil
		}

	}
	return 0, errors.New("data not found")
}

// TODO: custom this using crypto.Keccak256Hash instead of Hash ---- DONE(Verify success) =>>>>> NEXT: apply this to contract
func VerifyProofUsing(data []byte, proof *Proof, root []byte) (bool, error) {
	var dataHash []byte
	dataHash = crypto.Keccak256Hash(data).Bytes()
	index := proof.Index + (1 << uint(len(proof.Hashes)))
	//	if index >= uint64(len(proof.Hashes)) {
	//		return false, errors.New("invalid proof")
	//	}

	for _, hash := range proof.Hashes {
		if index%2 == 0 {
			dataHash = crypto.Keccak256Hash(append(dataHash, hash...)).Bytes()
		} else {
			dataHash = crypto.Keccak256Hash(append(hash, dataHash...)).Bytes()
		}
		index = index >> 1
	}
	return bytes.Equal(dataHash, root), nil
}
func (t *MerkleTree) Root() []byte {
	return t.nodes[1]
}
