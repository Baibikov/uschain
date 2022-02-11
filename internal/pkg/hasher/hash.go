package hasher

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

// PublicKey implement public key hashing
type PublicKey []byte

// Hash 2 factor hasher information
// first  - 256 hashing
// second -  ripemd 160
func (p PublicKey) Hash() ([]byte, error) {
	pubHash := sha256.Sum256(p)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		return nil, errors.Wrap(
			err,
			"write hasher to public key",
		)
	}

	return hasher.Sum(nil), nil
}

// Equal used to compare bytes
// Example:
// a == a -> true
// b == c -> false
// ...
func (p PublicKey) Equal(hash []byte) bool {
	return bytes.Compare(p, hash) == 0
}

func Sum256(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

func ToHex(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"xex typing value: %d",
			num,
		)
	}

	return buff.Bytes(), nil
}


func HashCmp(target *big.Int, hash []byte) bool {
	intHash := big.Int{}
	return intHash.SetBytes(hash).Cmp(target) == -1
}