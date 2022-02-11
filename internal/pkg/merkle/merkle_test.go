package merkle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MerkleTree(t *testing.T) {

	tests := []struct{
		Name string
		Args [][]byte
		Want string
	}{
		{
			Name: "a + b merkle hash",
			Args: [][]byte{
				[]byte("a"),
				[]byte("b"),
			},
			Want: "e5a01fee14e0ed5c48714f22180f25ad8365b53f9779f79dc4a3d7e93963f94a",
		},
		{
			Name: "a + b + c + c merkle hash",
			Args: [][]byte{
				[]byte("a"),
				[]byte("b"),
				[]byte("c"),
			},
			Want: "d31a37ef6ac14a2db1470c4316beb5592e6afd4465022339adafda76a18ffabe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tree, err := NewTree(tt.Args)
			require.NoError(t, err)
			sqp := fmt.Sprintf("%x", tree.TopHash())
			require.Equal(t, sqp, tt.Want)
		})
	}
}