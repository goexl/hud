package hud

import "net/http"

type lifecycle interface {
	// Initiate 初始化
	Initiate() (id string, url string, err error)

	// Abort 取消
	Abort(id string)

	// Complete 完成
	Complete(headers []http.Header)
}
