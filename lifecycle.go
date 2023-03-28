package hud

import (
	"github.com/gabriel-vasile/mimetype"
)

type lifecycle interface {
	// Initiate 初始化
	Initiate(parts int, start int, mime *mimetype.MIME) (urls []string, err error)

	// Abort 取消
	Abort() (err error)

	// Complete 完成
	Complete(parts []*Part) (err error)
}
