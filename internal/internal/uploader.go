package internal

type Uploader struct {
	worker Worker
}

func NewUploader(worker Worker) *Uploader {
	return &Uploader{
		worker: worker,
	}
}

func (u *Uploader) Do() error {
	return u.worker.Do()
}
