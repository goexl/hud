package param

import (
	"github.com/goexl/http"

	"github.com/goexl/log"
)

type Hud struct {
	Http   *http.Client
	Logger log.Logger
}

func NewHud() *Hud {
	return &Hud{
		Http:   http.New().Build(),
		Logger: log.New().Apply(),
	}
}
