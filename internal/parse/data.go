package parse

import (
	"log"
	"strconv"
	"strings"
)

type Data struct {
	Time  int
	Value string
}

func NewData(source string) *Data {
	strs := strings.Split(source, "\n")
	if len(strs) == 3 {
		tStr := strs[2][0:8]
		tNum, err := strconv.Atoi(tStr)
		if err != nil {
			log.Printf("“%s” 报错: %v", source, err)
		} else {
			return &Data{
				Time:  tNum,
				Value: strs[0],
			}
		}
	}
	return nil
}
