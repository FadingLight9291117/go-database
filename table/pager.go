package table

import (
	"com.fadinglight/db/BTree"
	"fmt"
	"os"
)

type Pager struct {
	Pages    [MAX_PAGE]*Page
	PageNums int
	File     *os.File // 持有一个file指针
	FileSize int
}

func NewPager(path string) *Pager { return new(Pager).init(path) }

func (p *Pager) init(path string) *Pager {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("file cannot open")
	}
	p.File = file
	fileInfo, _ := p.File.Stat()
	p.FileSize = int(fileInfo.Size())
	p.PageNums = int(fileInfo.Size() / BTree.PAGE_SIZE)

	if fileInfo.Size()%BTree.PAGE_SIZE != 0 {
		fmt.Printf("%s size is %d. Db file is not a whole number of pages. Corrupt file.\n", path, fileInfo.Size())
		os.Exit(-1)
	}

	return p
}

// GetPage nodeType: 当page不存在时，新建的node类型
func (p *Pager) GetPage(pageNum int, nodeType BTree.NodeType) *Page {
	if pageNum > MAX_PAGE {
		panic("Tried to fetch page number out of bounds.")
	}

	if p.Pages[pageNum] == nil {
		pageSize := BTree.PAGE_SIZE
		// Cache miss.Allocate memory and load from file.
		var node BTree.Node
		switch nodeType {
		case BTree.NODE_TYPE_INTERNAL:
			node = BTree.CreateInternalNode()
		case BTree.NODE_TYPE_LEAF:
			node = BTree.CreateLeafNode()
		}
		p.Pages[pageNum] = &Page{node}

		// 文件中的 `page` 总数
		filePagesNum := p.FileSize / pageSize
		if p.FileSize%pageSize != 0 {
			filePagesNum++
		}

		// page在文件已有数据中
		if pageNum < filePagesNum {
			b := make([]byte, pageSize)
			offset := int64(pageNum * pageSize)

			_, err := p.File.ReadAt(b, offset)
			if err != nil {
				return nil
			}
			node, err := BTree.DeserializeNode(b)
			if err != nil {
				return nil
			}
			pg := &Page{node}
			p.Pages[pageNum] = pg
			if pageNum >= p.PageNums {
				p.PageNums++
			}
		}
		if pageNum >= filePagesNum {
			p.PageNums++
		}
	}

	return p.Pages[pageNum]
}

// GetUnusedPageNum
// Until we start recycling free pages, new pages will always
// go onto the end of the database file
func (p *Pager) GetUnusedPageNum() int {
	return p.PageNums
}

func (p *Pager) FlushOnePage(pageIndex int) error {
	page := p.Pages[pageIndex]
	if page == nil {
		return nil
	}
	buf, err := BTree.SerializeNode(page.Node)
	if err != nil {
		return err
	}
	offset := int64(BTree.PAGE_SIZE * pageIndex)
	_, err = p.File.WriteAt(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pager) PrintTree(pageNum int, indentationLevel int) {
	indent := func(indentationLevel int) {
		for i := 0; i < indentationLevel; i++ {
			fmt.Print("  ")
		}
	}
	node := p.GetPage(pageNum, 0).Node

	switch node.(type) {
	case *BTree.LeafNode:
		leafNode := node.(*BTree.LeafNode)
		keyNums := leafNode.CellNums
		cells := leafNode.Cells[:keyNums]
		indent(indentationLevel)
		fmt.Printf("- leaf (size %d)", keyNums)
		if leafNode.IsNodeRoot() {
			fmt.Printf(" (Root)")
		}
		fmt.Println()
		for _, cell := range cells {
			indent(indentationLevel + 1)
			fmt.Printf("- %d\n", cell.Key)
		}
	case *BTree.InternalNode:
		internalNode := node.(*BTree.InternalNode)
		keyNums := internalNode.CellNums
		cells := internalNode.Cells[:keyNums]
		indent(indentationLevel)
		fmt.Printf("- internal (size %d)", keyNums)
		if internalNode.IsNodeRoot() {
			fmt.Printf(" (Root)")
		}
		fmt.Println()
		for _, cell := range cells {
			childIdx := cell.ChildPointer
			p.PrintTree(int(childIdx), indentationLevel+1)
			indent(indentationLevel + 1)
			fmt.Printf("- key %d\n", cell.Key)
		}
		rightChild := internalNode.RightChild
		p.PrintTree(int(rightChild), indentationLevel+1)
	default:
		return
	}
}
