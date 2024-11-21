package parse

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Bucket []*Data

func NewBucket(list []string) Bucket {
	var bucket Bucket
	for _, item := range list {
		data := NewData(item)
		if data != nil {
			bucket = append(bucket, data)
		}
	}
	return bucket
}

func (b Bucket) Len() int {
	return len(b)
}

func (b Bucket) Less(i, j int) bool {
	return b[i].Time < b[j].Time
}

func (b Bucket) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Bucket) Save(filePath string) {
	// 创建CSV文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 初始化CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"日期", "内容"}
	if err = writer.Write(header); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	// 写入记录
	for _, record := range b {
		row := []string{
			fmt.Sprintf("%v", record.Time),
			record.Value,
		}
		if err = writer.Write(row); err != nil {
			fmt.Println("Error writing row:", err)
			return
		}
	}
}
