package shift

import (
	"strings"
)

type Shift struct {
	lineTxt  []string
}

// set lineTxt in this class lineTxt
func (shift *Shift) SetLineTxt(lineTxt []string) {
	shift.lineTxt = lineTxt
}

// get lineTxt
func (shift *Shift) GetLineTxt() []string {
	return shift.lineTxt
}

// set kwicList in this class by lineTxt
func (shift *Shift) Kwic() {
	cnt := len(shift.lineTxt)
	// 获取每一个单词添加到words切片中
	for _, str := range shift.lineTxt {
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
			shift.lineTxt = append(shift.lineTxt, tStr)
			tStr = ""
			reverseCnt--
		}
		cnt--
		if cnt == 0 {
			break
		}
	}
}

func reverse(a []string, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}


