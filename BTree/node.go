package BTree

import (
	"bytes"
	"encoding/binary"
)

type Node interface {
	Serialize() ([]byte, error)
	Deserialize(b []byte) error
	IsNodeRoot() bool
	SetNodeRoot(isRoot bool)
	GetMaxKey() uint64
}

// CreateLeafNode 将node的cellNum置零
func CreateLeafNode() *LeafNode {
	node := &LeafNode{}
	node.Type = NODE_TYPE_LEAF
	node.IsRoot = false
	node.CellNums = 0
	return node
}

func CreateInternalNode() *InternalNode {
	node := &InternalNode{}
	node.Type = NODE_TYPE_INTERNAL
	node.IsRoot = false
	node.CellNums = 0
	return node
}

func DeserializeNode(b []byte) (Node, error) {
	var node Node

	commonNodeHeader := &CommonNodeHeader{}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, commonNodeHeader)
	if err != nil {
		return nil, err
	}
	switch commonNodeHeader.Type {
	case NODE_TYPE_INTERNAL:
		node = CreateInternalNode()
	case NODE_TYPE_LEAF:
		node = CreateLeafNode()
	}
	err = node.Deserialize(b)
	if err != nil {
		return nil, err
	}

	return node, nil
}
