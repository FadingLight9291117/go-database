package BTree

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

type NodeType = uint8

const (
	NODE_INTERNAL NodeType = iota
	NODE_LEAF
)

type CommonNodeHeader struct {
	Type      NodeType
	IsRoot    bool
	ParentPtr *CommonNodeHeader
}

type LeafNodeHeader struct {
	CommonNodeHeader
	CellNums uint32
}

type NodeCell struct {
	Key   uint32
	Value Row
}

type LeafNode struct {
	LeafNodeHeader
	Cells [LEAF_NODE_MAX_CELLS]NodeCell
}

/**
 * Common Node Header Layout (10 byte)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 byte)
 */
const (
	NODE_TYPE_SIZE          = int(unsafe.Sizeof(NODE_INTERNAL))
	NODE_TYPE_OFFSET        = int(0)
	NODE_IS_ROOT_SIZE       = int(unsafe.Sizeof(true))
	NODE_IS_ROOT_OFFSET     = int(NODE_TYPE_OFFSET + NODE_TYPE_SIZE)
	NODE_PARENTPTR_SIZE     = int(unsafe.Sizeof(&CommonNodeHeader{}))
	NODE_PARENTPTR_OFFSET   = int(NODE_IS_ROOT_OFFSET + NODE_IS_ROOT_SIZE)
	COMMON_NODE_HEADER_SIZE = int(NODE_TYPE_SIZE + NODE_IS_ROOT_SIZE + NODE_PARENTPTR_SIZE)
)

/**
 * Leaf Node Header Layout (14 byte)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 bytes) + cell_nums(4 bytes)
 */
const (
	LEAF_NODE_CELL_NUMS_SIZE   = int(unsafe.Sizeof(uint32(0)))
	LEAF_NODE_CELL_NUMS_OFFSET = int(COMMON_NODE_HEADER_SIZE)
	LEAF_NODE_HEADER_SIZE      = int(COMMON_NODE_HEADER_SIZE + LEAF_NODE_CELL_NUMS_SIZE)
)

/**
 *	Leaf Node Body Layout
 *  key(8 bytes) - value(164 bytes) (N)
 */
const (
	LEAF_NODE_KEY_SIZE        = int(unsafe.Sizeof(uint32(0)))
	LEAF_NODE_KEY_OFFSET      = int(0)
	LEAF_NODE_VALUE_SIZE      = int(ROW_SIZE)
	LEAF_NODE_VALUE_OFFSET    = int(LEAF_NODE_KEY_OFFSET + LEAF_NODE_KEY_SIZE)
	LEAF_NODE_CELL_SIZE       = int(LEAF_NODE_KEY_SIZE + LEAF_NODE_VALUE_SIZE)
	LEAF_NODE_SPACE_FOR_CELLS = int(PAGE_SIZE - LEAF_NODE_HEADER_SIZE)
	LEAF_NODE_MAX_CELLS       = int(LEAF_NODE_SPACE_FOR_CELLS / LEAF_NODE_CELL_SIZE)
)

func GetLeafNodeCellsNum(b *[]byte) (int, error) {
	var cellNum uint32
	buf := bytes.NewBuffer((*b)[LEAF_NODE_CELL_NUMS_OFFSET:])
	if err := binary.Read(buf, binary.BigEndian, &cellNum); err != nil {
		return -1, err
	}
	return int(cellNum), nil
}

func GetLeafNodeCell(b *[]byte, cellNum int) (*NodeCell, error) {
	buf := bytes.NewBuffer((*b)[LEAF_NODE_HEADER_SIZE : LEAF_NODE_HEADER_SIZE+cellNum*LEAF_NODE_CELL_SIZE])
	var cell *NodeCell
	err := binary.Read(buf, binary.BigEndian, cell)
	if err != nil {
		return nil, err
	}
	return cell, nil
}

func GetLeafNodeKey(b *[]byte, cellNum int) (uint32, error) {
	cell, err := GetLeafNodeCell(b, cellNum)
	if err != nil {
		return 0, err
	}
	return cell.Key, nil
}

func GetLeafNodeValue(b *[]byte, cellNum int) (*Row, error) {
	cell, err := GetLeafNodeCell(b, cellNum)
	if err != nil {
		return nil, err
	}
	return &cell.Value, nil
}

// InitLeafNode 将node的cellNum置零
func InitLeafNode(b *[]byte) error {
	buf := new(bytes.Buffer)
	err := binary.Read(buf, binary.BigEndian, 0)
	if err != nil {
		return err
	}
	copy((*b)[LEAF_NODE_CELL_NUMS_OFFSET:LEAF_NODE_CELL_NUMS_OFFSET+LEAF_NODE_CELL_NUMS_SIZE], buf.Bytes())
	return nil
}
