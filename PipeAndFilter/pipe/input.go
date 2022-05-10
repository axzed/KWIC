package pipe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Input struct {
	lineTxt []string
}

func (input *Input) GetLineTxt() []string {
	return input.lineTxt
}

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
	for {
		str, err := reader.ReadString('\n')
		re := regexp.MustCompile("\r\n")
		str = re.ReplaceAllString(str, "")
		input.lineTxt = append(input.lineTxt, str)
		if err == io.EOF {
			break
		}
	}	
}


