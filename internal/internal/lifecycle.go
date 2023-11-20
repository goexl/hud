package internal

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/goexl/hud/internal/bo"
)

type Lifecycle interface {
	// Initiate 初始化
	Initiate(parts int, start int, mime *mimetype.MIME) (urls []*bo.Url, err error)

	// Abort 取消
	Abort() (err error)

	// Complete 完成
	Complete(parts []*bo.Part) (err error)
}
