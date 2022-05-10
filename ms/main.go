package main

import "kiwc_xwc/ms/kwic"

func main() {
	fileName := "D:\\kwic_xwc\\ms\\chooseAndSort.txt"
	kwic.Input(fileName)
	kwic.Shift()
	kwic.Sort()
	kwic.Output("result.txt")
}
