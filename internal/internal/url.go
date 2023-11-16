package internal

import (
	"github.com/goexl/gox/http"
)

// Url 地址
type Url struct {
	// 连接方法
	Method http.Method `json:"method,omitempty"`
	// 目标地址
	Target string `json:"target,omitempty"`
}
