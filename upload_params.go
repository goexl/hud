package hud

import (
	"io"
	"os"
)

type uploadParams struct {
	target io.ReadSeeker
}

func newUploadParams() *uploadParams {
	return &uploadParams{}
}

func (up *uploadParams) bytes(params *multipartParams, part int64) (parts int, err error) {

}

func (up *uploadParams) parts(params *multipartParams) (parts int64, err error) {
	if size, se := up.size(); nil != se {
		err = se
	} else {
		parts = size / params.size.Byte()
	}

	return
}

func (up *uploadParams) size() (size int64, err error) {
	if bytes, bok := up.target.([]byte); bok {
		size = int64(len(bytes))
	} else if info, se := os.Stat(up.target.(string)); nil != se {
		err = se
	} else {
		size = info.Size()
	}

	return
}
