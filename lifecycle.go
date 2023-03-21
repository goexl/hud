package hud

import (
	"time"
)

type lifecycle interface {
	// Request 请求上传
	Request() (url string, expired time.Time, err error)

	// Initiate 初始化
	Initiate() (id string, err error)

	// Abort 取消
	Abort(id string)

	// Complete 完成
	Complete()
}
