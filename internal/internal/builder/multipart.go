package builder

import (
	"github.com/goexl/gox"
	"github.com/goexl/hud/internal/internal"
	"github.com/goexl/hud/internal/param"
	"github.com/goexl/hud/internal/worker"
)

type Multipart struct {
	builder *Upload
	params  *param.Hud
	upload  *param.Upload
	self    *param.Multipart
}

func NewMultipart(builder *Upload, params *param.Hud, upload *param.Upload) *Multipart {
	return &Multipart{
		builder: builder,
		params:  params,
		upload:  upload,
		self:    param.NewMultipart(),
	}
}

func (m *Multipart) Max(max int) (multipart *Multipart) {
	m.self.Max = max
	multipart = m

	return
}

func (m *Multipart) Size(size gox.Bytes) (multipart *Multipart) {
	m.self.Size = size
	multipart = m

	return
}

func (m *Multipart) Lifecycle(lifecycle internal.Lifecycle) (multipart *Multipart) {
	m.self.Lifecycle = lifecycle
	multipart = m

	return
}

func (m *Multipart) Build() (upload *Upload) {
	m.builder.worker = worker.NewWorkerMultipart(m.params, m.upload, m.self)
	upload = m.builder

	return
}
