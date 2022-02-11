package blockchain

import (
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"

	"uschain/internal/pkg/hasher"
)

const sumLength = 4

type TxOutput struct {
	Value      int

	publicKeyHash []byte
}

// initPublicKeyHash init transaction output public key hasher
func (tx *TxOutput) initPublicKeyHash(address string) error {
	ph, err := base58.Decode(address)
	if err != nil {
		return errors.Wrap(
			err,
			"public key decode",
		)
	}

	ph = ph[1 : len(ph)-sumLength]
	tx.publicKeyHash = ph

	return nil
}

// HashEqual compare public key hasher
func (tx *TxOutput) HashEqual(publicKeyHash []byte) bool {
	return hasher.PublicKey(tx.publicKeyHash).Equal(publicKeyHash)
}

// NewTxOutput output transaction constructor
func NewTxOutput(value int, address string) (*TxOutput, error) {
	output := TxOutput{
		Value: value,
	}

	err :=  output.initPublicKeyHash(address)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

type TxOutputs struct {
	Outputs []TxOutput
}


