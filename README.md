# video-ocr

通过ocr将东方财富的历史净值遍历出来

## 技术栈

必要：
- 安装 [tesseract-ocr/tesseract](https://github.com/tesseract-ocr/tesseract)
- go get github.com/otiai10/gosseract/v2

练习了多线程的使用方式，代码在`internal/ocr/pool.go`中，欢迎交流