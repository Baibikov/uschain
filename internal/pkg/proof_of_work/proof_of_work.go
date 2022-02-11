package proof_of_work

import (
	"bytes"
	"math"
	"math/big"

	"github.com/pkg/errors"

	"uschain/internal/pkg/hasher"
)

type Proffer interface {
	Validate(nonce int) (bool, error)
	Run() (int, []byte, error)
}


type proof struct {
	difficulty 			int
	target 				*big.Int
	transactionsHash 	[]byte
	prevHash			[]byte
}

func NewProf(difficulty int, prevHash, transactionsHash []byte) Proffer {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	return &proof{
		difficulty: difficulty,
		target: target,
		transactionsHash: transactionsHash,
		prevHash: prevHash,
	}
}

func (p proof) Validate(nonce int) (bool, error) {
	data, err := p.init(nonce)
	if err != nil {
		return false, err
	}

	return hasher.HashCmp(p.target, hasher.Sum256(data)), nil
}

func (p *proof) Run() (nonce int, hash []byte, err error) {
	for nonce < math.MaxInt {
		data := make([]byte, 0)
		data, err = p.init(nonce)
		if err != nil {
			return 0, nil, errors.Wrap(
				err,
				"init pow of proof",
			)
		}

		hash = hasher.Sum256(data)
		if hasher.HashCmp(p.target, hash) {
			break
		}

		nonce++
	}

	return nonce, hash, nil
}

func (p *proof) init(nonce int) ([]byte, error) {
	difHash, err := hasher.ToHex(int64(p.difficulty))
	if err != nil {
		return nil, err
	}

	nonceHash, err := hasher.ToHex(int64(nonce))
	if err != nil {
		return nil, err
	}

	return bytes.Join([][]byte{
		p.prevHash,
		p.transactionsHash,
		difHash,
		nonceHash,
	}, []byte{}), nil

}
