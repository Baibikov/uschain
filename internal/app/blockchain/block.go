package blockchain

import (
	"time"

	"uschain/internal/pkg/proof_of_work"
)

type Block struct {
	Timestamp   	int64
	Hash         	[]byte
	Transactions 	[]*Transaction
	PrevHash     	[]byte
	Nonce        	int
	Height       	int
}

const difficulty = 12

func CreateBlock(txs []*Transaction, prevHash []byte, height int) (*Block, error) {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Transactions: txs,
		PrevHash: prevHash,
		Height: height,
	}

	topHash, err := Transactions(txs).MerkleTopHash()
	if err != nil {
		return nil, err
	}

	block.Nonce, block.Hash, err = proof_of_work.NewProf(difficulty, prevHash, topHash).Run()
	if err != nil {
		return nil, err
	}

	return block, nil
}

func Genesis(coinbase *Transaction) (*Block, error) {
	return CreateBlock([]*Transaction{coinbase}, []byte{}, 0)
}