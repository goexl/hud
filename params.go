package hud

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type params struct {
	timeout time.Duration
	http    *resty.Client
}

func newParams() *params {
	return &params{
		timeout: 5 * time.Second,
	}
}
