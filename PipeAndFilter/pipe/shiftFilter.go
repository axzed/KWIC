package pipe

import (
	"strings"
)

type ShiftFilter struct{}

func NewShiftFilter() *ShiftFilter {
	return &ShiftFilter{}
}

// 实现Filter接口
func (sft *ShiftFilter) Filter(selectLine []string) (shiftLine []string) {
	cnt := len(selectLine)
	// 获取每一个单词添加到words切片中
	for _, str := range selectLine {
		temp := strings.Split(str, " ")
		n := len(temp)
		reverseCnt := len(temp) - 1
		idx := 1
		for reverseCnt != 0 {
			reverse(temp, 0, idx-1)
			reverse(temp, idx, n-1)
			reverse(temp, 0, n-1)
			var tStr string
			for _, word := range temp {
				tStr += word
				tStr += " "

			}
			shiftLine = append(shiftLine, tStr)
			tStr = ""
			reverseCnt--
		}
		cnt--
		if cnt == 0 {
			break
		}
	}
	return shiftLine
}

func reverse(a []string, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
