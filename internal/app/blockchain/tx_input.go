package blockchain

import (
	"uschain/internal/pkg/hasher"
)

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    hasher.PublicKey
}

