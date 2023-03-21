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

func (ub *uploadBuilder) Multipart() *multipartBuilder {
	return newMultipartBuilder(ub)
}

func (ub *uploadBuilder) Build() *uploader {
	return newUploader(ub.worker)
}
