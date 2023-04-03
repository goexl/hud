package hud

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/http"
)

var _ worker = (*workerMultipart)(nil)

type workerMultipart struct {
	params *params
	upload *uploadParams
	self   *multipartParams
	parts  []*Part
}

func newWorkerMultipart(params *params, upload *uploadParams, self *multipartParams) *workerMultipart {
	return &workerMultipart{
		params: params,
		upload: upload,
		self:   self,
	}
}

func (wm *workerMultipart) do() (err error) {
	if parts, pe := wm.upload.parts(wm.self); nil != pe {
		err = pe
	} else if mime, me := wm.upload.mime(); nil != me {
		err = me
	} else if urls, ie := wm.self.lifecycle.Initiate(parts, wm.self.start, mime); nil != ie { // 第一步，初始化上分片上传
		err = ie
	} else {
		wm.parts = make([]*Part, 0, parts)
		wm.uploads(urls, parts, &err)
	}

	return
}

func (wm *workerMultipart) uploads(urls []*Url, parts int, err *error) {
	wg := new(sync.WaitGroup)
	wg.Add(parts)
	for part := wm.self.start; part < parts+wm.self.start; part++ {
		go wm.part(urls[part-wm.self.start], part, wg, err) // 第二步，按分片上传对应的文件切片
	}
	wg.Wait()

	// 第三步，如果没有错误，完成上传
	if nil == *err {
		*err = wm.self.lifecycle.Complete(wm.parts)
	}
}

func (wm *workerMultipart) part(url *Url, part int, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if bytes, be := wm.upload.bytes(wm.self, part); nil != be {
		*err = be
	} else if rsp, pe := wm.send(url, bytes); nil != pe {
		*err = pe
	} else if rsp.IsError() {
		*err = exc.NewException(rsp.StatusCode(), "上传到服务器出错", field.New("response", rsp.String()))
	} else {
		_part := new(Part)
		_part.Number = int32(part)
		_part.Header = rsp.Header()
		_part.Size = gox.Size(len(bytes))
		wm.parts = append(wm.parts, _part)
	}
}

func (wm *workerMultipart) send(url *Url, bytes []byte) (rsp *resty.Response, err error) {
	req := wm.params.http.R().SetBody(bytes)
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
