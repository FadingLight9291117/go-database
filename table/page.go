package table

import (
	"bytes"
	"com.fadinglight/db/BTree"
	"encoding/binary"
)

type Page struct {
	BTree.LeafNode
}

func (p *Page) Serialize() ([]byte, error) {
	bs := make([]byte, BTree.PAGE_SIZE)
	buf := bytes.NewBuffer(bs)
	if err := binary.Write(buf, binary.BigEndian, p); err != nil {
		return nil, err
	}
	copy(bs, buf.Bytes())

	return bs, nil
}

func DeserializePage(bs []byte) (*Page, error) {
	page := new(Page)
	buf := bytes.NewBuffer(bs[:BTree.PAGE_SIZE-BTree.LEAF_NODE_SPACE_FOR_CELLS])
	if err := binary.Read(buf, binary.BigEndian, &page); err != nil {
		return nil, err
	}

	return page, nil
}
