package hud

import (
	"github.com/goexl/gox"
)

type multipartParams struct {
	lifecycle lifecycle
	max       int
	size      gox.Size
}

func newMultipartParams() *multipartParams {
	return &multipartParams{
		max:  10000,
		size: gox.SizeMB,
	}
}
