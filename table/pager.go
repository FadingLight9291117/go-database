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

//GetPage FIXME: only return a leaf node
func (p *Pager) GetPage(pageNum int) *Page {
	if pageNum > MAX_PAGE {
		panic("Tried to fetch page number out of bounds.")
	}

	if p.Pages[pageNum] == nil {
		pageSize := BTree.PAGE_SIZE
		// Cache miss.Allocate memory and load from file.
		p.Pages[pageNum] = &Page{BTree.CreateLeafNode()}

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
			pg := &Page{BTree.CreateLeafNode()}
			err = pg.Deserialize(b)
			if err != nil {
				return nil
			}
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
	buf, err := page.Serialize()
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
