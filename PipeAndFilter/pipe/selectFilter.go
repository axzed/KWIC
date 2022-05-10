package pipe

// 选择过滤器,选出10个句子
type SelectFilter struct {}

// 实例化选择过滤器
func NewSelectFilter() *SelectFilter {
	return &SelectFilter{}
}

// 实现Filter接口
func (sf *SelectFilter) Filter (lineTxt []string) (selectLine []string) {
	// 选择前10个句子
	for i := 0; i < 10; i ++ {
		selectLine = append(selectLine, lineTxt[i])
	}
	return 
}

