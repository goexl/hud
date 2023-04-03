package hud

import (
	"net/http"

	"github.com/goexl/gox"
)

// Part 分片
type Part struct {
	// 分片编号
	Number int32 `json:"number,omitempty"`
	// 头
	Header http.Header `json:"header,omitempty"`
	// 文件大小
	Size gox.Size `json:"size,omitempty"`
}
