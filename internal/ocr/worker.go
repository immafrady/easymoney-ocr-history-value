package ocr

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

type Worker struct {
	id     int
	client *gosseract.Client
	count  int
}

func NewWorker(id int) *Worker {
	w := &Worker{id: id}
	w.client = gosseract.NewClient()
	w.client.SetLanguage("chi_sim")
	return w
}

func (w *Worker) Exec(path string, log chan<- string) (string, error) {
	defer func() {
		w.count++ // 统计工作量
	}()

	log <- fmt.Sprintf("%v号: 处理 %s", w.id, path)
	err := w.client.SetImage(path)
	if err != nil {
		log <- fmt.Sprintf("%v号: 处理 %s失败， %v", w.id, path, err)
	}

	return w.client.Text()
}
