package blockchain

import (
	"crypto/rand"
	"fmt"

	"github.com/pkg/errors"

	"uschain/internal/pkg/hasher"
	"uschain/internal/pkg/merkle"
	"uschain/internal/pkg/ser"
)

const (
	baseTxOutputValue = 20
)

type Transaction struct {
	ID []byte

	Outputs []TxOutput
	Inputs  []TxInput
}

// Hash make hash of transaction
func (tx *Transaction) Hash() ([]byte, error) {
	cop := *tx
	tx.ID = nil

	buff, err := ser.Serialize(cop)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"transaction make hash from id %x",
			tx.ID,
		)
	}

	return hasher.Sum256(buff), nil
}

// NewTransaction make a new transaction
func NewTransaction(out []TxOutput, inp []TxInput) (*Transaction, error) {
	tx := Transaction{
		Outputs: out,
		Inputs: inp,
	}

	id, err := tx.Hash()
	if err != nil {
		return nil, err
	}
	tx.ID = id

	return &tx, err
}

// CoinBaseTx make init transaction for used by create blockchain
func CoinBaseTx(to, data string) (*Transaction, error) {
	if data == "" {
		dt, err := txRandomData()
		if err != nil {
			return nil, err
		}
		data = dt
	}

	txOut, err := NewTxOutput(baseTxOutputValue, to)
	if err != nil {
		return nil, err
	}

	txIn := TxInput{
		ID: []byte{},
		Out: -1,
		PubKey: []byte(data),
	}

	tx, err := NewTransaction([]TxOutput{*txOut}, []TxInput{txIn})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func txRandomData() (string, error) {
	randData := make([]byte, 24)
	_, err := rand.Read(randData)
	if err != nil {
		return "", errors.Wrap(
			err,
			"rand data where empty information",
		)
	}

	return fmt.Sprintf("%x", randData), nil
}

type Transactions []*Transaction

func (tx Transactions) MerkleTopHash() ([]byte, error) {
	transactions := make([][]byte, 0)
	for _, t := range tx {
		st, err := ser.Serialize(t)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, st)
	}

	t, err := merkle.NewTree(transactions)
	if err != nil {
		return nil, err
	}

	return t.TopHash(), nil
}
