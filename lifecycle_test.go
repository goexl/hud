package hud_test

import (
	"context"

	"github.com/gabriel-vasile/mimetype"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/hud"
	"go.sichuancredit.cn/protocol/storage"
	"go.sichuancredit.cn/protocol/storage/core"
	"go.sichuancredit.cn/protocol/storage/protocol"
)

type lifecycle struct {
	client storage.RpcClient
	id     int64
}

func newLifecycle(client storage.RpcClient) *lifecycle {
	return &lifecycle{
		client: client,
	}
}

func (l *lifecycle) Initiate(parts int, start int, mime *mimetype.MIME) (urls []string, err error) {
	req := new(protocol.InitiateReq)
	req.Bucket = "test"
	req.ContentType = mime.String()
	req.Parts = int32(parts)
	req.Start = int32(start)
	if rsp, err := l.client.Initiate(context.Background(), req); nil == err {
		l.id = rsp.Id
		urls = rsp.Urls
	}

	return
}

func (l *lifecycle) Abort(_ string) (err error) {
	return
}

func (l *lifecycle) Complete(parts []*hud.Part) (err error) {
	req := new(protocol.CompleteReq)
	req.Id = l.id
	req.Parts = make([]*core.Part, 0, len(parts))
	for _, _part := range parts {
		part := new(core.Part)
		part.Number = _part.Number
		part.Etag = _part.Header.Get("ETag")
		req.Parts = append(req.Parts, part)
	}
	if rsp, ce := l.client.Complete(context.Background(), req); nil != ce {
		err = ce
	} else if 0 == rsp.Id {
		err = exc.NewField("完成上传出错", field.New("parts", parts))
	}

	return
}
