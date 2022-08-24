package table

import (
	"os"
)

type Pager struct {
	Pages    [MAX_PAGE]*Page
	File     *os.File // 持有一个file指针
	FileSize int
}

func OpenPager(path string) *Pager { return new(Pager).init(path) }

func (p *Pager) init(path string) *Pager {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("file cannot open")
	}
	p.File = file
	fileInfo, _ := p.File.Stat()
	p.FileSize = int(fileInfo.Size())
	return p
}

func (p *Pager) GetPage(pageNum int) *Page {
	if pageNum > MAX_PAGE {
		panic("Tried to fetch page number out of bounds.")
	}

	if p.Pages[pageNum] == nil {
		pageSize := PAGE_SIZE
		// Cache miss.Allocate memory and load from file.
		p.Pages[pageNum] = &Page{}

		// 文件中的 `page` 总数
		filePagesNum := p.FileSize / pageSize
		if p.FileSize%pageSize != 0 {
			filePagesNum++
		}

		// page在文件已有数据中
		if pageNum < int(filePagesNum) {
			b := make([]byte, pageSize)
			offset := int64(pageNum * pageSize)

			if pageNum == int(filePagesNum)-1 {
				Len := p.FileSize % pageSize
				_, err := p.File.ReadAt(b[:Len], offset) // `b`的长度不可以大于`File`的长度
				if err != nil {
					return nil
				}
			} else {
				_, err := p.File.ReadAt(b, offset) // `b`的长度不可以大于`File`的长度
				if err != nil {
					return nil
				}
			}

			pg, err := DeserializePage(b)
			if err != nil {
				return nil
			}
			p.Pages[pageNum] = pg
		}
	}

	return p.Pages[pageNum]
}

func (p *Pager) FlushOnePage(pageIndex int, rowNum int) error {
	page := p.Pages[pageIndex]
	buf, err := page.Serialize()
	buf = buf[:rowNum*ROW_SIZE]
	if err != nil {
		return err
	}
	offset := int64(PAGE_SIZE * pageIndex)
	_, err = p.File.WriteAt(buf, offset)
	if err != nil {
		return err
	}
	err = p.File.Close()
	if err != nil {
		return err
	}
	return nil
}
