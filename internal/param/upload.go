package param

import (
	"os"

	"github.com/gabriel-vasile/mimetype"
)

type Upload struct {
	Target any
	File   *os.File
	Cap    int64
}

func NewUpload() *Upload {
	return &Upload{}
}

func (u *Upload) Bytes(params *Multipart, part int) (bytes []byte, err error) {
	if filepath, fok := u.Target.(string); fok && nil == u.File {
		u.File, err = os.Open(filepath)
	}
	if nil != err {
		return
	}

	size := 0
	offset := params.Offset(part)
	bytes = make([]byte, params.Cap(part, u.Cap))
	switch _target := u.Target.(type) {
	case []byte:
		bytes = _target[offset : int(offset)+len(bytes)]
		size = len(bytes)
	case string:
		size, err = u.File.ReadAt(bytes, offset)
	}
	if nil == err {
		bytes = bytes[:size]
	}

	return
}

func (u *Upload) Parts(params *Multipart) (parts int, err error) {
	if size, se := u.Size(); nil != se {
		err = se
	} else {
		parts = params.Parts(size)
	}

	return
}

func (u *Upload) Size() (size int64, err error) {
	if bytes, bok := u.Target.([]byte); bok {
		size = int64(len(bytes))
	} else if info, se := os.Stat(u.Target.(string)); nil != se {
		err = se
	} else {
		size = info.Size()
	}
	if nil == err {
		u.Cap = size
	}

	return
}

func (u *Upload) Mime() (mime *mimetype.MIME, err error) {
	switch target := u.Target.(type) {
	case []byte:
		mime = mimetype.Detect(target)
	case string:
		mime, err = mimetype.DetectFile(target)
	}

	return
}
