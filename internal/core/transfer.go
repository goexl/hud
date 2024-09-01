package core

import (
	"github.com/goexl/gox"
	"github.com/goexl/hud/internal/internal/builder"
	"github.com/goexl/hud/internal/param"
)

// Transfer 传输器
type Transfer struct {
	params *param.Hud
	_      gox.Pointerized
}

func NewTransfer(params *param.Hud) *Transfer {
	return &Transfer{
		params: params,
	}
}

func (t *Transfer) Upload() *builder.Upload {
	return builder.NewUpload(t.params)
}
