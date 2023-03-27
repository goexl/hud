package hud

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Http(http *resty.Client) *builder {
	b.params.http = http

	return b
}

func (b *builder) Logger(logger simaqian.Logger) *builder {
	b.params.logger = logger

	return b
}

func (b *builder) Build() *Transfer {
	return newTransfer(b.params)
}
