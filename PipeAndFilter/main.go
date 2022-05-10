package main

import (
	"fmt"
	"kwic/PipeAndFilter/pipe"
)

func main() {
	// 从文件中读入文本
	var input = new(pipe.Input)
	input.Input("D:\\kwic_xwc\\PipeAndFilter\\chooseAndSort.txt")

	// 过滤出选择的式子
	selector := pipe.NewSelectFilter()
	selectLine := selector.Filter(input.GetLineTxt())
	// test to print selectLine
	// for i, v := range selectLine {
	// 	fmt.Println(i, v)
	// }

	fmt.Println("--------------------------------------------------")

	// 过滤出位移
	shiftor := pipe.NewShiftFilter()
	shiftLine := shiftor.Filter(selectLine)
	// for i, v := range shiftLine {
	// 	fmt.Println(i, v)
	// }

	// 过滤出kwic序列
	sortor := pipe.NewSortAlphFilter()
	kwicLine := sortor.Filter(shiftLine)
	// test to print kwicLine
	// for i, v := range kwicLine {
	// 	fmt.Println(i, v)
	// }

	// 写入文件
	var output = new(pipe.Output)
	output.SetKwicList(kwicLine)
	output.SetNewFileName("result03.txt")
	output.Write()
}
