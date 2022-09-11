package table

import (
	"com.fadinglight/db/BTree"
	"fmt"
	"os"
)

type Pager struct {
	Pages    [MAX_PAGE]*Page
	PageNum  int
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
	p.PageNum = int(fileInfo.Size() / BTree.PAGE_SIZE)

	if fileInfo.Size()%BTree.PAGE_SIZE != 0 {
		fmt.Println("Db file is not a whole number of pages. Corrupt file.")
		os.Exit(-1)
	}

	return p
}

func (p *Pager) GetPage(pageNum int) *Page {
	if pageNum > MAX_PAGE {
		panic("Tried to fetch page number out of bounds.")
	}

	if p.Pages[pageNum] == nil {
		pageSize := BTree.PAGE_SIZE
		// Cache miss.Allocate memory and load from file.
		p.Pages[pageNum] = &Page{*BTree.InitLeafNode()}

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

			pg, err := DeserializePage(b)
			if err != nil {
				return nil
			}
			p.Pages[pageNum] = pg
			if pageNum >= p.PageNum {
				p.PageNum++
			}
		}
		if pageNum >= filePagesNum {
			p.PageNum++
		}
	}

	return p.Pages[pageNum]
}

func (p *Pager) FlushOnePage(pageIndex int) error {
	page := p.Pages[pageIndex]
	buf, err := page.Serialize()
	//buf = buf[:rowNum*ROW_SIZE]
	if err != nil {
		return err
	}
	offset := int64(BTree.PAGE_SIZE * pageIndex)
	_, err = p.File.WriteAt(buf, offset)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
