package main

import (
	"fmt"
	"os"
)

func getFileName(i int) string {
	imagePath := fmt.Sprintf("source/output_%03d.png", i)
	if _, err := os.Stat(imagePath); err != nil {
		return ""
	}
	return imagePath
}
