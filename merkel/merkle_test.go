package merkel

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
	}

	// Level 1
	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	// Level 2
	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	// Level 3
	n7 := NewMerkleNode(n5, n6, nil)

	fmt.Println(hex.EncodeToString(n7.Data))

	root := NewMerkleTree(data)
	fmt.Println(hex.EncodeToString(root.RootNode.Data))
	//assert.Equal(t, hex.EncodeToString(n7.Data), hex.EncodeToString(root.RootNode.Data), "Root hash is correct")
}
