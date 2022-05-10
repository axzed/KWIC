package pipe

import (
	"bufio"
	"fmt"
	"os"
)

type Output struct {
	kwicList    []string
	newFileName string
}

func (output *Output) SetKwicList(kwicList []string) {
	output.kwicList = kwicList
}

func (output *Output) SetNewFileName(newFileName string) {
	output.newFileName = newFileName
}

func (output *Output) Write() {
	file, err := os.OpenFile(output.newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	// 及时关闭file句柄
	defer file.Close()
	// 写入文件,使用带缓存的*Write,只是写入缓存
	write := bufio.NewWriter(file)
	for _, v := range output.kwicList {
		// 写入的时候在每个句尾加上换行符
		v += "\r\n"
		write.WriteString(v)
	}
	// Flush将缓存的文件真正写入到文件
	write.Flush()
	fmt.Println("写入成功!")
}
