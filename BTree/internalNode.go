package BTree

import "unsafe"

type Node interface {
}

type InternalNodeHeader struct {
	CommonNodeHeader
	CellNums   uint32
	RightChild *Node
}

type InternalNodeCell struct {
	Key   uint32
	Value *Node
}

type InternalNode struct {
	InternalNodeHeader
	Cells []InternalNodeCell
}

/**
 * Internal Node Header Layout
 */
const (
	INTERNAL_NODE_CELL_NUMS_SIZE   = int(unsafe.Sizeof(uint32(0)))
	INTERNAL_NODE_CELL_NUMS_OFFSET = int(unsafe.Offsetof(COMMON_NODE_HEADER_SIZE))
	INTERNAL_NODE_HEADER_SIZE      = int(unsafe.Sizeof(COMMON_NODE_HEADER_SIZE + INTERNAL_NODE_CELL_NUMS_SIZE))
)

/**
 * Internal Node Body Layout
 */
const (
	INTERNAL_NODE_KEY_SIZE   = int(unsafe.Sizeof(uint32(0)))
	INTERNAL_NODE_VALUE_SIZE = int(unsafe.Sizeof(unsafe.Sizeof(&LeafNode{})))
)
