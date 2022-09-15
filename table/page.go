package table

import (
	"com.fadinglight/db/BTree"
)

type Page struct {
	BTree.LeafNode
}

func (p *Page) Serialize() ([]byte, error) {
	b, err := BTree.SerializeLeafNode(&p.LeafNode)
	return b, err
}

func DeserializePage(bs []byte) (*Page, error) {
	node, err := BTree.DeSerializeLeafNode(bs)
	return &Page{*node}, err
}

func (p *Page) IsNodeRoot() bool {
	return p.IsRoot
}

func (p *Page) CopyFrom(op *Page) {
	p.LeafNode = op.LeafNode
}
