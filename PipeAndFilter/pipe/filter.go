package pipe

// 过滤器接口
type Filter interface {
	Filter([]string) []string
}