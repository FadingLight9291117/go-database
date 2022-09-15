package BTree

import "unsafe"

/**
 * Internal Node Header Layout (14 bytes)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 bytes) + cell_nums(4 bytes) + right_child(8 bytes)
 */
const (
	INTERNAL_NODE_CELL_NUMS_SIZE     = int(unsafe.Sizeof(uint64(0)))
	INTERNAL_NODE_CELL_NUMS_OFFSET   = int(COMMON_NODE_HEADER_SIZE)
	INTERNAL_NODE_RIGHT_CHILD_SIZE   = int(unsafe.Sizeof(uint64(0)))
	INTERNAL_NODE_RIGHT_CHILD_OFFSET = int(INTERNAL_NODE_CELL_NUMS_OFFSET + INTERNAL_NODE_CELL_NUMS_SIZE)
	INTERNAL_NODE_HEADER_SIZE        = int(COMMON_NODE_HEADER_SIZE + INTERNAL_NODE_CELL_NUMS_SIZE + INTERNAL_NODE_RIGHT_CHILD_SIZE)
)

/**
 * Internal Node Body Layout
 * key(8 bytes) - value(? bytes) (N)
 */
const (
	INTERNAL_NODE_KEY_SIZE        = int(unsafe.Sizeof(uint64(0)))
	INTERNAL_NODE_CHILD_SIZE      = int(unsafe.Sizeof(uint64(0)))
	INTERNAL_NODE_CELL_SIZE       = int(INTERNAL_NODE_KEY_SIZE + INTERNAL_NODE_CHILD_SIZE)
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

// InternalNodeHeader
/*
 * internal node 的内部节点指针总是比键多一个
 * 我们把这个额外的子指针放在标头中
 */
type InternalNodeHeader struct {
	CommonNodeHeader
	CellNums   uint64
	RightChild unsafe.Pointer
}
type InternalNodeCell struct {
	ChildPointer unsafe.Pointer
	Key          uint64
}
type InternalNode struct {
	InternalNodeHeader
	Cells [INTERNAL_NODE_MAX_CELLS]InternalNodeCell
}
