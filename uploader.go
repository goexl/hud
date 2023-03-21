package hud

type uploader struct {
	worker worker
}

func newUploader(worker worker) *uploader {
	return &uploader{
		worker: worker,
	}
}

func (u *uploader) Do() error {
	return u.worker.do()
}
