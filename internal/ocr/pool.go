package ocr

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Pool struct {
	workers map[int]*Worker
	free    chan int
	pending chan string
	result  chan string
	status  bool
	logs    chan string
}

func NewPool(l int) *Pool {
	p := &Pool{
		workers: make(map[int]*Worker),
		free:    make(chan int, l),
		pending: make(chan string, l),
		result:  make(chan string),
	}

	for i := 1; i <= l; i++ {
		p.workers[i] = NewWorker(i)
		p.free <- i // 将空闲的worker id推入
	}
	return p
}

func (p *Pool) listen() {
	for {
		path, ok := <-p.pending
		if !ok {
			count := len(p.workers)
			for {
				id := <-p.free
				p.logs <- fmt.Sprintf("%v号机归位, 干活：%v", id, p.workers[id].count)

				count--
				if count == 0 { // 所有空闲worker归位之后，等于结束
					close(p.free)
					close(p.result)
					return
				}
			}
		} else {
			id := <-p.free

			go func() {
				text, err := p.workers[id].Exec(path, p.logs)
				if err != nil {
					log.Println(err)
				}
				p.free <- id
				p.result <- text
			}()
		}
	}
}

func (p *Pool) GetResult(fileFormat string) []string {
	p.logs = make(chan string) // 每次开始前都确保有日志
	defer func() {
		close(p.logs)
	}()
	go p.printLogs()
	go p.listen()
	go func() {
		i := 1
		for {
			imagePath := fmt.Sprintf(fileFormat, i)
			if _, err := os.Stat(imagePath); os.IsNotExist(err) {
				p.logs <- "图片加载完成"
				break
			}
			p.logs <- fmt.Sprintf("已添加: %v, %s", i, imagePath)
			p.pending <- imagePath
			i++
		}
		close(p.pending)
	}()

	var ret []string
	for {
		result, ok := <-p.result
		if !ok {
			return ret
		}
		ret = append(ret, result)
	}
}

func (p *Pool) printLogs() {
	for l := range p.logs {
		t := time.Now()
		fmt.Printf("%s %v\n", t.Format("2006-01-02 15:04:05"), l)
	}
}
