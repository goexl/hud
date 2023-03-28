package hud

import (
	"github.com/goexl/gox"
)

type multipartParams struct {
	lifecycle lifecycle
	max       int
	size      gox.Size
	start     int
}

func newMultipartParams() *multipartParams {
	return &multipartParams{
		max:   10000,
		size:  5 * gox.SizeMB,
		start: 1,
	}
}

func (mp *multipartParams) parts(size int64) int {
	return int(size/mp.size.Byte()) + mp.start
}

func (mp *multipartParams) offset(part int) int64 {
	return mp.size.Byte() * int64(part-mp.start)
}

func (mp *multipartParams) cap(part int, size int64) (cap int) {
	total := int64(part) * mp.size.Byte()
	if total > size {
		cap = int(size - mp.offset(part))
	} else {
		cap = int(mp.size.Byte())
	}

	return
}
