package builder

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/hud/internal/core"
	"github.com/goexl/hud/internal/param"
	"github.com/goexl/log"
)

type Hud struct {
	params *param.Hud
}

func NewHud() *Hud {
	return &Hud{
		params: param.NewHud(),
	}
}

func (h *Hud) Http(http *resty.Client) *Hud {
	h.params.Http = http

	return h
}

func (h *Hud) Logger(logger log.Logger) *Hud {
	h.params.Logger = logger

	return h
}

func (h *Hud) Build() *core.Transfer {
	return core.NewTransfer(h.params)
}
