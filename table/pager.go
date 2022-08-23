package table

import (
	"os"
)

type Pager struct {
	Pages [MAX_PAGE]*Page
	File  *os.File // 持有一个file指针
}

func OpenPager(path string) *Pager { return new(Pager).init(path) }

func (p *Pager) init(path string) *Pager {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("file cannot open")
	}
	p.File = file
	return p
}

// Len FIXME: 这里没有存入文件
func (p *Pager) Len() int {
	for i := range p.Pages {
		if p.Pages[i] == nil {
			return i
		}
	}
	return 0
}

func (p *Pager) GetPage(pageNum int) *Page {
	if pageNum > MAX_PAGE {
		panic("Tried to fetch page number out of bounds.")
	}

	if p.Pages[pageNum] == nil {
		// Cache miss.Allocate memory and load from file.
		p.Pages[pageNum] = &Page{}

		fileStat, _ := p.File.Stat()
		// 文件中的 `page` 总数
		filePagesNum := fileStat.Size() / PAGE_SIZE // FIXME
		if fileStat.Size()%PAGE_SIZE != 0 {
			filePagesNum++
		}

		//// page在文件已有数据中
		//if pageNum < int(filePagesNum) {
		//	// TODO: 读取文件内容, [pageNum*PAGE_SIZE, pageNum*PAGE_SIZE + PAGE_SIZE)
		//	b := make([]byte, PAGE_SIZE)
		//	_, err := p.File.ReadAt(b, int64(PAGE_SIZE*pageNum))
		//	if err != nil {
		//		return nil
		//	}
		//	pg, err := DeserializePage(b)
		//	if err != nil {
		//		return nil
		//	}
		//	p.Pages[pageNum] = pg
		//}
	}

	return p.Pages[pageNum]
}
