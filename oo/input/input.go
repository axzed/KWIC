package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// a class to input
type Input struct {
	lineTxt []string
}

// get lineTxt
func (input *Input) GetLineTxt() []string {
	return input.lineTxt
}

// open file & set lineTxt
func (input *Input) Input(fileName string) {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // 关闭文件
	// 新建一个缓冲区,把内容放在缓冲区
	reader := bufio.NewReader(file)
	// 循环按行读取,取出10个句子
	cnt := 10
	for cnt != 0 {
		str, err := reader.ReadString('\n')
		// windows下用regexp包编译"\r\n", 因为只是txt文件再windows下的换行,一定要记住用re编译然后替换源字符串
		re := regexp.MustCompile("\r\n")
		str = re.ReplaceAllString(str, "")
		input.lineTxt = append(input.lineTxt, str)
		cnt--
		if err == io.EOF {
			break
		}
	}
}
