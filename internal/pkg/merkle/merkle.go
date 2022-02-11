package merkle

import (
	"github.com/pkg/errors"

	"uschain/internal/pkg/hasher"
)

type Node struct {
	hash []byte

	left *Node
	right *Node
}

func (n *Node) IsEmpty() bool {
	return n == nil
}

// makeHash has make hash for node hashed data
func (n *Node) makeHash(data []byte) {
	if n.left.IsEmpty() && n.right.IsEmpty() {
		n.hash = hasher.Sum256(data)
		return
	}

	n.hash = hasher.Sum256(append(
		n.left.hash,
		n.right.hash...,
	))
}

// NewNode make new node from tree
func NewNode(left, right *Node, data []byte) *Node {
	node := &Node{
		left: left,
		right: right,
	}
	node.makeHash(data)

	return node
}

type Tree struct {
	root *Node
}

// TopHash get the upper hash from tree
func (t *Tree) TopHash() []byte {
	return t.root.hash
}

// NewTree make merkle tree
// 			[hash(a+b)] 			- root
//   	    /          \
//      [hash(a)]	  [hash(b)] 	- first level
//       /               \
//   [hash(l1)]		  [hash(l2)]	- second level.
//      |				  |
//      l1				  l2        - data blocks
func NewTree(data [][]byte) (*Tree, error) {
	if len(data) == 0 {
		return nil, errors.New("merkle has not nodes")
	}

	// make base nodes - L
	nodes := make([]Node, 0)
	for _, d := range data {
		nodes = append(nodes, *NewNode(nil, nil, d))
	}

	// node addition
	if len(nodes) % 2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	// we pass to all nodes
	// and make upper hash
	for len(nodes) > 1 {
		level := make([]Node, 0)
		for i := 0; i < len(nodes); i+=2 {
			level = append(level, *NewNode(&nodes[i], &nodes[i+1], nil))
		}
		nodes = level
	}

	return &Tree{
		root: &nodes[0],
	}, nil
}
