package hud

type multipartBuilder struct {
	builder *uploadBuilder
	params  *multipartParams
}

func newMultipartBuilder(builder *uploadBuilder) *multipartBuilder {
	return &multipartBuilder{
		params: newMultipartParams(),
	}
}

func (mb *multipartBuilder) Build() *uploadBuilder {
	mb.builder.worker = newWorkerMultipart(mb.params)

	return mb.builder
}
