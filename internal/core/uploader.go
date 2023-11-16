package core

import (
	"github.com/goexl/hud/internal/internal"
)

type Uploader struct {
	worker internal.Worker
}

func NewUploader(worker internal.Worker) *Uploader {
	return &Uploader{
		worker: worker,
	}
}

func (u *Uploader) Do() error {
	return u.worker.Do()
}
