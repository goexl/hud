package hud

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type params struct {
	http   *resty.Client
	logger simaqian.Logger
}

func newParams() *params {
	return &params{
		http:   resty.New(),
		logger: simaqian.Default(),
	}
}
