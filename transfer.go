package hud

import (
	"github.com/goexl/gox"
)

// Transfer 传输器
type Transfer struct {
	params *params
	_      gox.CannotCopy
}

func newTransfer(params *params) *Transfer {
	return &Transfer{
		params: params,
	}
}

func (t *Transfer) Upload() *uploadBuilder {
	return newUploadBuilder(t.params)
}
