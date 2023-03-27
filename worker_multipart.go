package hud

import (
	"net/http"
	"sync"
)

var _ worker = (*workerMultipart)(nil)

type workerMultipart struct {
	params  *params
	upload  *uploadParams
	self    *multipartParams
	headers []http.Header
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
	} else if id, url, ie := wm.self.lifecycle.Initiate(); nil != ie { // 第一步，初始化上分片上传
		err = ie
	} else {
		wm.uploads(url, id, parts, &err)
	}

	return
}

func (wm *workerMultipart) uploads(url string, id string, parts int, err *error) {
	wg := new(sync.WaitGroup)
	wg.Add(parts)
	for part := 0; part < parts; part++ {
		go wm.part(url, part, wg, err) // 第二步，按分片上传对应的文件切片
	}
	wg.Wait()

	// 第三步，如果没有错误，完成上传
	if nil == *err {
		wm.self.lifecycle.Complete(id, wm.headers)
	}
}

func (wm *workerMultipart) part(url string, part int, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if bytes, be := wm.upload.bytes(wm.self, part); nil != be {
		*err = be
	} else if rsp, pe := wm.params.http.R().SetBody(bytes).Put(url); nil != pe {
		*err = pe
	} else {
		wm.headers = append(wm.headers, rsp.Header())
	}
}
