package hud

import (
	"github.com/goexl/gox"
)

type multipartBuilder struct {
	builder *uploadBuilder
	params  *params
	upload  *uploadParams
	self    *multipartParams
}

func newMultipartBuilder(builder *uploadBuilder, params *params, upload *uploadParams) *multipartBuilder {
	return &multipartBuilder{
		builder: builder,
		params:  params,
		upload:  upload,
		self:    newMultipartParams(),
	}
}

func (mb *multipartBuilder) Max(max int) *multipartBuilder {
	mb.self.max = max

	return mb
}

func (mb *multipartBuilder) Size(size gox.Size) *multipartBuilder {
	mb.self.size = size

	return mb
}

func (mb *multipartBuilder) Lifecycle(lifecycle lifecycle) *multipartBuilder {
	mb.self.lifecycle = lifecycle

	return mb
}

func (mb *multipartBuilder) Build() *uploadBuilder {
	mb.builder.worker = newWorkerMultipart(mb.params, mb.upload, mb.self)

	return mb.builder
}
