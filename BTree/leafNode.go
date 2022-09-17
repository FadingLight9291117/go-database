package BTree

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

/**
 * LeafNode Header Layout (26 bytes)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 bytes) - nextLeaf(8 bytes) - cellNums(8 bytes)
 */
const (
	LEAF_NODE_NEXT_LEAF_SIZE   = int(unsafe.Sizeof(uint64(0)))
	LEAF_NODE_CELL_NUMS_SIZE   = int(unsafe.Sizeof(uint64(0)))
	LEAF_NODE_CELL_NUMS_OFFSET = int(COMMON_NODE_HEADER_SIZE)
	LEAF_NODE_HEADER_SIZE      = int(COMMON_NODE_HEADER_SIZE + LEAF_NODE_NEXT_LEAF_SIZE + LEAF_NODE_CELL_NUMS_SIZE)
)

/**
 *	LeafNode Body Layout
 *  key(8 bytes) - value(164 bytes) (N)
 */
const (
	LEAF_NODE_KEY_SIZE        = int(unsafe.Sizeof(uint64(0)))
	LEAF_NODE_KEY_OFFSET      = int(0)
	LEAF_NODE_VALUE_SIZE      = int(ROW_SIZE)
	LEAF_NODE_VALUE_OFFSET    = int(LEAF_NODE_KEY_OFFSET + LEAF_NODE_KEY_SIZE)
	LEAF_NODE_CELL_SIZE       = int(LEAF_NODE_KEY_SIZE + LEAF_NODE_VALUE_SIZE)
	LEAF_NODE_SPACE_FOR_CELLS = int(PAGE_SIZE - LEAF_NODE_HEADER_SIZE)
	LEAF_NODE_MAX_CELLS       = int(LEAF_NODE_SPACE_FOR_CELLS / LEAF_NODE_CELL_SIZE)
)

/**
 * leaf node other size
 */
const (
	LEAF_NODE_SIZE        = int(LEAF_NODE_HEADER_SIZE + LEAF_NODE_CELL_SIZE*LEAF_NODE_MAX_CELLS)
	LEAF_NODE_WASTED_SIZE = int(PAGE_SIZE - LEAF_NODE_SIZE)
)

type LeafNodeHeader struct {
	CommonNodeHeader
	NextLeaf uint64 // 0 represents no sibling
	CellNums uint64
}

type LeafNodeCell struct {
	Key   uint64
	Value Row
}

type LeafNode struct {
	LeafNodeHeader
	Cells [LEAF_NODE_MAX_CELLS]LeafNodeCell
}

func (node *LeafNode) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, node)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (node *LeafNode) Deserialize(b []byte) error {
	buf := bytes.NewReader(b[:LEAF_NODE_SIZE])
	err := binary.Read(buf, binary.BigEndian, node)
	if err != nil {
		return err
	}
	return nil
}

func (node *LeafNode) IsNodeRoot() bool {
	return node.IsRoot
}

func (node *LeafNode) SetNodeRoot(isRoot bool) {
	node.IsRoot = isRoot
}

func (node *LeafNode) GetMaxKey() uint64 {
	return node.Cells[node.CellNums-1].Key
}

func (node *LeafNode) GetNextLeafNodeNum() int {
	return int(node.NextLeaf)
}
