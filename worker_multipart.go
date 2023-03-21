package hud

var _ worker = (*workerMultipart)(nil)

type workerMultipart struct {
	params *params
	upload *uploadParams
	self   *multipartParams
}

func newWorkerMultipart(params *params, upload *uploadParams, self *multipartParams) *workerMultipart {
	return &workerMultipart{
		params: params,
		upload: upload,
		self:   self,
	}
}

func (wm *workerMultipart) do() (err error) {
	return
}
