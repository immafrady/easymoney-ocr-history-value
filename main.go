package main

import (
	"fmt"
	"github.com/immafrady/video-ocr/internal/ocr"
	"github.com/immafrady/video-ocr/internal/parse"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func main() {
	t1 := time.Now()
	defer func() {
		t2 := time.Now()
		fmt.Printf("总计用时：%vms", t2.Sub(t1).Milliseconds())
	}()

	p := ocr.NewPool(10)

	base, _ := os.Getwd()
	result := p.GetResult(filepath.Join(base, "source/output_%03d.png"))

	bucket := parse.NewBucket(result)
	sort.Sort(bucket)
	bucket.Save("output.csv")
}
