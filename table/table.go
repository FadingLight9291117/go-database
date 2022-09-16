package table

import (
	"com.fadinglight/db/BTree"
)

type Table struct {
	Pager       *Pager
	RootPageNum int
}

func NewTable(filename string) *Table { return new(Table).init(filename) }

func (t *Table) init(filename string) *Table {
	pager := NewPager(filename)
	t.Pager = pager
	t.RootPageNum = 0
	if pager.PageNums == 0 {
		// New database file. Initialize page 0 as leaf node.
		t.Pager.GetPage(0, 0).SetNodeRoot(true)
	}

	return t
}

func (t *Table) Select() []*BTree.Row {
	rows := make([]*BTree.Row, 0)

	pageNum := t.Pager.FileSize / BTree.PAGE_SIZE

	for i := 0; i < pageNum; i++ {
		page := t.Pager.GetPage(i, 0)
		if node, ok := page.Node.(*BTree.LeafNode); ok {
			for j := 0; j < int(node.CellNums); j++ {
				rows = append(rows, &node.Cells[j].Value)
			}
		}

	}

	return rows
}

func (t *Table) Close() error {
	pager := t.Pager
	for i := 0; i < pager.PageNums; i++ {
		err := pager.FlushOnePage(i)
		if err != nil {
			return err
		}
	}
	err := t.Pager.File.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t *Table) CreateNewRoot(rightChildPageNum int) {
	/*
		Handle splitting thr root
		Old root copied to new Pag, becomes left child.
		Address of right child passed in.
		Re-initialize root page to contain the new root node.
		New root node points to the two child
	*/
	root := t.Pager.GetPage(t.RootPageNum, 0)
	//rightChild := t.Pager.GetPage(rightChildPageNum)
	leftChildNum := t.Pager.GetUnusedPageNum()
	leftChild := t.Pager.GetPage(leftChildNum, 0)
	*leftChild = *root
	leftChild.SetNodeRoot(false)
	// new root
	*root = Page{BTree.CreateInternalNode()}
	root.SetNodeRoot(true)
	if rootNode, ok := root.Node.(*BTree.InternalNode); ok {
		rootNode.CellNums = 1
		rootNode.Cells[0].ChildPointer = uint64(leftChildNum)
		rootNode.Cells[0].Key = leftChild.GetMaxKey()
		rootNode.RightChild = uint64(rightChildPageNum)
	} else {
		panic("error: not a internal node")
	}
}
