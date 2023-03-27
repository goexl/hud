package hud

type uploadBuilder struct {
	params *params
	self   *uploadParams
	worker worker
}

func newUploadBuilder(params *params) *uploadBuilder {
	return &uploadBuilder{
		params: params,
		self:   newUploadParams(),
	}
}

func (ub *uploadBuilder) Bytes(bytes []byte) *uploadBuilder {
	ub.self.target = bytes

	return ub
}

func (ub *uploadBuilder) File(path string) *uploadBuilder {
	ub.self.target = path

	return ub
}

func (ub *uploadBuilder) Multipart() *multipartBuilder {
	return newMultipartBuilder(ub, ub.params, ub.self)
}

func (ub *uploadBuilder) Build() *uploader {
	return newUploader(ub.worker)
}
