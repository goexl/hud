package hud

import (
	"github.com/goexl/hud/internal/builder"
)

// New 创建构造器
func New() *builder.Hud {
	return builder.NewHud()
}
