package BTree

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

/**
 * InternalNode Header Layout (14 bytes)
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
 * InternalNode Body Layout
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
	INTERNAL_NODE_SIZE        = int(INTERNAL_NODE_HEADER_SIZE + INTERNAL_NODE_CELL_SIZE*INTERNAL_NODE_MAX_CELLS)
	INTERNAL_NODE_WASTED_SIZE = int(PAGE_SIZE - INTERNAL_NODE_SIZE)
)

// InternalNodeHeader
/*
 * internal node 的内部节点指针总是比键多一个
 * 我们把这个额外的子指针放在标头中
 */
type InternalNodeHeader struct {
	CommonNodeHeader
	CellNums   uint64
	RightChild uint64
}
type InternalNodeCell struct {
	ChildPointer uint64
	Key          uint64
}
type InternalNode struct {
	InternalNodeHeader
	Cells [INTERNAL_NODE_MAX_CELLS]InternalNodeCell
}

func (node *InternalNode) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, node)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (node *InternalNode) Deserialize(b []byte) error {
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, node)
	if err != nil {
		return err
	}
	return nil
}

func (node *InternalNode) IsNodeRoot() bool {
	return node.IsRoot
}

func (node *InternalNode) SetNodeRoot(isRoot bool) {
	node.IsRoot = isRoot
}

func (node *InternalNode) GetMaxKey() uint64 {
	return node.Cells[node.CellNums-1].Key
}

func (node *InternalNode) GetParentNum() int {
	return int(node.ParentNum)
}

func (node *InternalNode) SetParentNum(num int) {
	node.ParentNum = uint64(num)
}

func (node *InternalNode) UpdateKey(oldKey uint64, newKey uint64) {
	cellIndex := node.FindChildIdx(oldKey)
	node.Cells[cellIndex].Key = newKey
}

func (node *InternalNode) FindChildIdx(key uint64) int {
	minIndex := 0
	maxIndex := int(node.CellNums)

	for minIndex < maxIndex {
		midIndex := (minIndex + maxIndex) / 2
		if node.Cells[midIndex].Key == key {
			minIndex = midIndex
			break
		}
		if node.Cells[minIndex].Key > key {
			maxIndex = minIndex
		} else {
			minIndex = midIndex + 1
		}
	}
	return minIndex
}
