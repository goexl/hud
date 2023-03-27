package hud

var _ = New

// New 创建构造器
func New() *builder {
	return newBuilder()
}
