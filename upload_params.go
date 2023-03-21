package hud

import (
	"os"
)

type uploadParams struct {
	target any
	file *os.File
}

func newUploadParams() *uploadParams {
	return &uploadParams{}
}

func (up *uploadParams) bytes(params *multipartParams, part int) (bytes []byte, err error) {
	offset:=params.size.Byte()*int64(part)
	if filepath,fok:=up.target.(string);fok&&nil==up.file{
		up.file,err = os.Open(filepath)
	}
	if nil!=err{
		return
	}

	size:=0
	switch _target:=up.target.(type) {
	case []byte:
		size=len(_target)
	case string:
		size,err=up.file.ReadAt(bytes, offset)
	}
	if nil==err{
		bytes=bytes[:size]
	}

	return
}

func (up *uploadParams) parts(params *multipartParams) (parts int, err error) {
	if size, se := up.size(); nil != se {
		err = se
	} else {
		parts = int(size / params.size.Byte())
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
