package sortAlph

import (
	"fmt"
	"sort"
)

type SortAlph struct {
	kwicList []string
}

func (sortAlph *SortAlph) SetKwicList(kwicList []string) {
	sortAlph.kwicList = kwicList
}

func (sortAlph *SortAlph) GetKwicList() []string {
	return sortAlph.kwicList
}

func (sortAlph *SortAlph) Test() {
	for _, v := range sortAlph.kwicList {
		fmt.Println(v)
	}
}

// a function to sortwords && reverse it
func (sortAlph *SortAlph) Sort() {
	// 字母表顺序排序
	sort.Strings(sortAlph.kwicList)
	n := len(sortAlph.kwicList)
	// 逆序
	reverse(sortAlph.kwicList, 0, n-1)
}

func reverse(a []string, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
