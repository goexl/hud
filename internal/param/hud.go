package param

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/log"
)

type Hud struct {
	Http   *resty.Client
	Logger log.Logger
}

func NewHud() *Hud {
	return &Hud{
		Http:   resty.New(),
		Logger: log.New().Apply(),
	}
}
