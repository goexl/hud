package worker

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/http"
	"github.com/goexl/hud/internal/bo"
	"github.com/goexl/hud/internal/internal"
	"github.com/goexl/hud/internal/param"
)

var _ internal.Worker = (*Multipart)(nil)

type Multipart struct {
	params *param.Hud
	upload *param.Upload
	self   *param.Multipart
	parts  []*bo.Part
}

func NewWorkerMultipart(params *param.Hud, upload *param.Upload, self *param.Multipart) *Multipart {
	return &Multipart{
		params: params,
		upload: upload,
		self:   self,
	}
}

func (m *Multipart) Do() (err error) {
	if parts, pe := m.upload.Parts(m.self); nil != pe {
		err = pe
	} else if mime, me := m.upload.Mime(); nil != me {
		err = me
	} else if urls, ie := m.self.Lifecycle.Initiate(parts, m.self.Start, mime); nil != ie { // 第一步，初始化上分片上传
		err = ie
	} else {
		m.parts = make([]*bo.Part, 0, parts)
		m.uploads(urls, parts, &err)
	}

	return
}

func (m *Multipart) uploads(urls []*bo.Url, parts int, err *error) {
	wg := new(sync.WaitGroup)
	wg.Add(parts)
	for part := m.self.Start; part < parts+m.self.Start; part++ {
		go m.part(urls[part-m.self.Start], part, wg, err) // 第二步，按分片上传对应的文件切片
	}
	wg.Wait()

	// 第三步，如果没有错误，完成上传
	if nil == *err {
		*err = m.self.Lifecycle.Complete(m.parts)
	}
}

func (m *Multipart) part(url *bo.Url, part int, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if bytes, be := m.upload.Bytes(m.self, part); nil != be {
		*err = be
	} else if rsp, pe := m.send(url, bytes); nil != pe {
		*err = pe
	} else if rsp.IsError() {
		*err = exc.NewException(rsp.StatusCode(), "上传到服务器出错", field.New("response", rsp.String()))
	} else {
		_part := new(bo.Part)
		_part.Number = int32(part)
		_part.Header = rsp.Header()
		_part.Size = gox.Size(len(bytes))
		m.parts = append(m.parts, _part)
	}
}

func (m *Multipart) send(url *bo.Url, bytes []byte) (rsp *resty.Response, err error) {
	req := m.params.Http.R().SetBody(bytes)
	switch url.Method {
	case http.MethodPut:
		rsp, err = req.Put(url.Target)
	case http.MethodPost:
		rsp, err = req.Post(url.Target)
	case http.MethodGet:
		rsp, err = req.Get(url.Target)
	case http.MethodDelete:
		rsp, err = req.Delete(url.Target)
	}

	return
}
