package hud

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/log"
)

type params struct {
	http   *resty.Client
	logger log.Logger
}

func newParams() *params {
	return &params{
		http:   resty.New(),
		logger: log.New().Apply(),
	}
}
