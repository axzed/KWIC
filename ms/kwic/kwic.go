package kwic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

// lineTxt: 存放句子的切片
var lineTxt []string

// Input: a function to open a file and add in lineTxt to process
func Input(fileName string) {
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
		// fmt.Println(str)
		lineTxt = append(lineTxt, str)
		cnt--
		if err == io.EOF {
			break
		}
	}
}

// Ouput: a function to write result in an new file
func Output(newFileName string) {
	file, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	// 及时关闭file句柄
	defer file.Close()
	// 写入文件,使用带缓存的*Write,只是写入缓存
	write := bufio.NewWriter(file)
	for _, v := range lineTxt {
		// 写入的时候在每个句尾加上换行符
		v += "\r\n"
		write.WriteString(v)
	}
	// Flush将缓存的文件真正写入到文件
	write.Flush()
}

// a function to shift words
func Shift() {
	cnt := len(lineTxt)
	// 获取每一个单词添加到words切片中
	for _, str := range lineTxt {
		temp := strings.Split(str, " ")
		n := len(temp)
		reverseCnt := len(temp) - 1
		idx := 1
		for reverseCnt != 0 {
			Reverse(temp, 0, idx-1)
			Reverse(temp, idx, n-1)
			Reverse(temp, 0, n-1)
			// fmt.Println(temp)
			var tStr string
			for _, word := range temp {
				tStr += word
				tStr += " "

			}
			lineTxt = append(lineTxt, tStr)
			tStr = ""
			reverseCnt--
		}
		cnt--
		if cnt == 0 {
			break
		}
	}
}

// a function to sortwords && reverse it
func Sort() {
	// 字母表顺序排序
	sort.Strings(lineTxt)
	n := len(lineTxt)
	// 逆序
	Reverse(lineTxt, 0, n-1)
}

func Reverse(a []string, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// func Test() {
// 	for _, v := range lineTxt {
// 		fmt.Println(v)
// 	}
// }

// func TestStrings() {
// 	var a = [...]string{"abc", "efg", "b", "A", "eee"}
// 	sort.Strings(a[:])
// 	fmt.Println(a)
// }
