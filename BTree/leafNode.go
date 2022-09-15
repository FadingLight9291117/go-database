package BTree

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

/**
 * Leaf Node Header Layout (10 byte)
 * type(1 byte) - isRoot(1 byte) - parentPtr(8 bytes) + cell_nums(8 bytes)
 */
const (
	LEAF_NODE_CELL_NUMS_SIZE   = int(unsafe.Sizeof(uint64(0)))
	LEAF_NODE_CELL_NUMS_OFFSET = int(COMMON_NODE_HEADER_SIZE)
	LEAF_NODE_HEADER_SIZE      = int(COMMON_NODE_HEADER_SIZE + LEAF_NODE_CELL_NUMS_SIZE)
)

/**
 *	Leaf Node Body Layout
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
	LEAF_NODE_SIZE       = int(LEAF_NODE_HEADER_SIZE + LEAF_NODE_CELL_SIZE*LEAF_NODE_MAX_CELLS)
	LEAF_NODE_BLANK_SIZE = int(PAGE_SIZE - LEAF_NODE_SIZE)
)

type LeafNodeHeader struct {
	CommonNodeHeader
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

func GetLeafNodeCell(b *[]byte, cellNum int) (*LeafNodeCell, error) {
	buf := bytes.NewBuffer((*b)[LEAF_NODE_HEADER_SIZE : LEAF_NODE_HEADER_SIZE+cellNum*LEAF_NODE_CELL_SIZE])
	var cell *LeafNodeCell
	err := binary.Read(buf, binary.BigEndian, cell)
	if err != nil {
		return nil, err
	}
	return cell, nil
}

func GetLeafNodeKey(b *[]byte, cellNum int) (uint64, error) {
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

// CreateLeafNode 将node的cellNum置零
func CreateLeafNode() *LeafNode {
	node := &LeafNode{
		LeafNodeHeader: LeafNodeHeader{
			CommonNodeHeader: CommonNodeHeader{Type: NODE_TYPE_LEAF},
		},
	}
	return node
}

func SerializeLeafNode(node *LeafNode) ([]byte, error) {
	buf := new(bytes.Buffer)
	data := []any{node.Type, node.IsRoot, uint64(0), node.CellNums, node.Cells}
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			return nil, err
		}
	}
	space := make([]byte, LEAF_NODE_BLANK_SIZE)
	buf.Write(space)
	return buf.Bytes(), nil
}

func (node *LeafNode) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	data := []any{node.Type, node.IsRoot, uint64(0), node.CellNums, node.Cells}
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			return nil, err
		}
	}
	space := make([]byte, LEAF_NODE_BLANK_SIZE)
	buf.Write(space)
	return buf.Bytes(), nil
}

func DeSerializeLeafNode(b []byte) (*LeafNode, error) {
	buf := bytes.NewReader(b[:LEAF_NODE_SIZE])

	var data struct {
		Type      NodeType
		IsRoot    bool
		ParentPtr uint64
		CellNums  uint64
		Cells     [LEAF_NODE_MAX_CELLS]LeafNodeCell
	}
	err := binary.Read(buf, binary.BigEndian, &data)
	if err != nil {
		return nil, err
	}

	leafNode := &LeafNode{
		LeafNodeHeader: LeafNodeHeader{
			CommonNodeHeader: CommonNodeHeader{
				Type:      data.Type,
				IsRoot:    data.IsRoot,
				ParentPtr: nil,
			},
			CellNums: data.CellNums,
		},
		Cells: data.Cells,
	}
	return leafNode, nil
}

func (node *LeafNode) Deserialize(b []byte) error {
	buf := bytes.NewReader(b[:LEAF_NODE_SIZE])

	var data struct {
		Type      NodeType
		IsRoot    bool
		ParentPtr uint64
		CellNums  uint64
		Cells     [LEAF_NODE_MAX_CELLS]LeafNodeCell
	}
	err := binary.Read(buf, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	node.LeafNodeHeader.CommonNodeHeader.Type = data.Type
	node.LeafNodeHeader.CommonNodeHeader.IsRoot = data.IsRoot
	node.LeafNodeHeader.CommonNodeHeader.ParentPtr = nil
	node.LeafNodeHeader.CellNums = data.CellNums
	node.Cells = data.Cells
	return nil
}

func (node *LeafNode) IsRoot() {
	return
}
