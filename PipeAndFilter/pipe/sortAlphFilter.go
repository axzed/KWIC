package pipe

import "sort"

type SortAlphFilter struct{}

func NewSortAlphFilter() *SortAlphFilter {
	return &SortAlphFilter{}
}

func (saf *SortAlphFilter) Filter(shiftLine []string) (kwicLine []string) {
	sort.Strings(shiftLine)
	n := len(shiftLine)
	reverse(shiftLine, 0, n-1)
	for i := 0; i < n; i ++ {
		kwicLine = append(kwicLine, shiftLine[i])
	}
	return
}