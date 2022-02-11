package hasher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPublicKey_Equal(t *testing.T) {
	test := []struct{
		name string
		hash1 []byte
		hash2 []byte
		want bool
	}{
		{
			name: "is-simple-equeal",
			hash1: []byte("ce97657a211b6a1e23e3d39d586c0ab481f9a84f"),
			hash2: []byte("ce97657a211b6a1e23e3d39d586c0ab481f9a84f"),
			want: true,
		},
		{
			name: "is-simple-false",
			hash1: []byte("ce97657a211b6a1e23e3d39d586c0ab481f9v313f"),
			hash2: []byte("ce97657a211b6a1e23e3d39d586c0ab481f9a84f"),
			want: false,
		},
	}

	for _, tt := range test {
		require.Equal(t, PublicKey(tt.hash1).Equal(tt.hash2), tt.want)
	}
}
