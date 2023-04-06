package hud

import (
	"os"

	"github.com/gabriel-vasile/mimetype"
)

type uploadParams struct {
	target any
	file   *os.File
	cap    int64
}

func newUploadParams() *uploadParams {
	return &uploadParams{}
}

func (up *uploadParams) bytes(params *multipartParams, part int) (bytes []byte, err error) {
	if filepath, fok := up.target.(string); fok && nil == up.file {
		up.file, err = os.Open(filepath)
	}
	if nil != err {
		return
	}

	size := 0
	offset := params.offset(part)
	bytes = make([]byte, params.cap(part, up.cap))
	switch _target := up.target.(type) {
	case []byte:
		bytes = _target[offset : int(offset)+len(bytes)]
		size = len(bytes)
	case string:
		size, err = up.file.ReadAt(bytes, offset)
	}
	if nil == err {
		bytes = bytes[:size]
	}

	return
}

func (up *uploadParams) parts(params *multipartParams) (parts int, err error) {
	if size, se := up.size(); nil != se {
		err = se
	} else {
		parts = params.parts(size)
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
	if nil == err {
		up.cap = size
	}

	return
}

func (up *uploadParams) mime() (mime *mimetype.MIME, err error) {
	switch target := up.target.(type) {
	case []byte:
		mime = mimetype.Detect(target)
	case string:
		mime, err = mimetype.DetectFile(target)
	}

	return
}
