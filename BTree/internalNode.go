package BTree

import "unsafe"

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
 * Internal Node Header Layout (14 bytes)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 bytes) + cell_nums(4 bytes)
 */
const (
	INTERNAL_NODE_CELL_NUMS_SIZE   = int(unsafe.Sizeof(uint32(0)))
	INTERNAL_NODE_CELL_NUMS_OFFSET = int(unsafe.Sizeof(COMMON_NODE_HEADER_SIZE))
	INTERNAL_NODE_HEADER_SIZE      = int(unsafe.Sizeof(COMMON_NODE_HEADER_SIZE + INTERNAL_NODE_CELL_NUMS_SIZE))
)

/**
 * Internal Node Body Layout
 * key(8 bytes) - value(? bytes) (N)
 */
const (
	INTERNAL_NODE_KEY_SIZE        = int(unsafe.Sizeof(uint32(0)))
	INTERNAL_NODE_VALUE_SIZE      = int(unsafe.Sizeof(unsafe.Sizeof(uint32(0))))
	INTERNAL_NODE_CELL_SIZE       = int(INTERNAL_NODE_KEY_SIZE + INTERNAL_NODE_VALUE_SIZE)
	INTERNAL_NODE_SPACE_FOR_CELLS = int(PAGE_SIZE - INTERNAL_NODE_HEADER_SIZE)
	INTERNAL_NODE_MAX_CELLS       = int(INTERNAL_NODE_SPACE_FOR_CELLS / INTERNAL_NODE_CELL_SIZE)
)

/**
 * internal node other size
 */
const (
	INTERNAL_NODE_SIZE       = int(INTERNAL_NODE_HEADER_SIZE + INTERNAL_NODE_CELL_SIZE*INTERNAL_NODE_MAX_CELLS)
	INTERNAL_NODE_BLANK_SIZE = int(PAGE_SIZE - INTERNAL_NODE_SIZE)
)
