package hud_test

import (
	"context"

	"github.com/gabriel-vasile/mimetype"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/http"
	"github.com/goexl/hud"
	"go.sichuancredit.cn/protocol/storage/core"
	"go.sichuancredit.cn/protocol/storage/file"
)

type lifecycle struct {
	ctx      context.Context
	client   file.RpcClient
	app      int64
	filename string
	id       *int64
}

func newLifecycle(ctx context.Context, client file.RpcClient, app int64, filename string, id *int64) *lifecycle {
	return &lifecycle{
		ctx:      ctx,
		client:   client,
		app:      app,
		filename: filename,
		id:       id,
	}
}

func (l *lifecycle) Initiate(parts int, start int, mime *mimetype.MIME) (urls []*hud.Url, err error) {
	req := new(file.InitiateReq)
	req.App = l.app
	req.Mime = mime.String()
	req.Name = l.filename
	req.Parts = int32(parts)
	req.Start = int32(start)
	if rsp, ie := l.client.Initiate(l.ctx, req); nil != ie {
		err = ie
	} else {
		*l.id = rsp.Id
		urls = l.convert(rsp.Urls)
	}

	return
}

func (l *lifecycle) Abort() (err error) {
	req := new(file.AbortReq)
	req.Id = *l.id
	if rsp, ae := l.client.Abort(l.ctx, req); nil != ae {
		err = ae
	} else if 0 == rsp.Id {
		err = exc.NewField("取消上传出错", field.New("id", req.Id))
	}

	return
}

func (l *lifecycle) Complete(parts []*hud.Part) (err error) {
	req := new(file.CompleteReq)
	req.Id = *l.id
	req.Parts = make([]*core.Part, 0, len(parts))
	for _, _part := range parts {
		part := new(core.Part)
		part.Number = _part.Number
		part.Etag = _part.Header.Get("ETag")
		part.Size = _part.Size.Byte()
		req.Parts = append(req.Parts, part)
	}
	if rsp, ce := l.client.Complete(l.ctx, req); nil != ce {
		err = ce
	} else if 0 == rsp.Id {
		err = exc.NewField("完成上传出错", field.New("parts", parts))
	}

	return
}

func (l *lifecycle) convert(urls []*core.Url) (to []*hud.Url) {
	to = make([]*hud.Url, 0, len(urls))
	for _, from := range urls {
		url := new(hud.Url)
		url.Method = http.ParseMethod(from.Method)
		url.Target = from.Target
		to = append(to, url)
	}

	return
}
