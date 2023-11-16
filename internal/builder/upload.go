package builder

import (
	"os"

	"github.com/goexl/hud/internal/core"
	"github.com/goexl/hud/internal/internal"
	"github.com/goexl/hud/internal/param"
)

type Upload struct {
	params *param.Hud
	self   *param.Upload
	worker internal.Worker
}

func NewUpload(params *param.Hud) *Upload {
	return &Upload{
		params: params,
		self:   param.NewUpload(),
	}
}

func (u *Upload) Bytes(bytes []byte) (upload *Upload) {
	u.self.Target = bytes
	upload = u

	return
}

func (u *Upload) Filepath(path string) (upload *Upload) {
	u.self.Target = path
	upload = u

	return
}

func (u *Upload) File(file *os.File) (upload *Upload) {
	u.self.Target = file
	upload = u

	return
}

func (u *Upload) Multipart() *Multipart {
	return NewMultipart(u, u.params, u.self)
}

func (u *Upload) Build() *core.Uploader {
	return core.NewUploader(u.worker)
}
