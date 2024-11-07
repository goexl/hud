package param

import (
	"github.com/goexl/gox"
	"github.com/goexl/hud/internal/internal"
)

type Multipart struct {
	Lifecycle internal.Lifecycle
	Max       int
	Size      gox.Bytes
	Start     int
}

func NewMultipart() *Multipart {
	return &Multipart{
		Max:   10000,
		Size:  64 * gox.BytesMB,
		Start: 1,
	}
}

func (m *Multipart) Parts(size int64) int {
	return int(size/m.Size.Byte()) + m.Start
}

func (m *Multipart) Offset(part int) int64 {
	return m.Size.Byte() * int64(part-m.Start)
}

func (m *Multipart) Cap(part int, size int64) (cap int) {
	total := int64(part) * m.Size.Byte()
	if total > size {
		cap = int(size - m.Offset(part))
	} else {
		cap = int(m.Size.Byte())
	}

	return
}
