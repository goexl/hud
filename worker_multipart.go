package hud

import "sync"

var _ worker = (*workerMultipart)(nil)

type workerMultipart struct {
	params *params
	upload *uploadParams
	self   *multipartParams
}

func newWorkerMultipart(params *params, upload *uploadParams, self *multipartParams) *workerMultipart {
	return &workerMultipart{
		params: params,
		upload: upload,
		self:   self,
	}
}

func (wm *workerMultipart) do() (err error) {
	if parts,pe:=wm.upload.parts(wm.self);nil!=pe{
		err=pe
	}else if url,re:=wm.self.lifecycle.Request();nil!=re{ // 第一步，请求上传地址
	err=re
}else if id, ie:=wm.self.lifecycle.Initiate();nil!=ie{ // 第二步，初始化上分片上传
	err=ie
}else {
wm.uploads(url, id, parts, &err)
}

	return
}

func (wm *workerMultipart) uploads(url string, id string, parts int,err *error) {
	wg:=new(sync.WaitGroup)
	wg.Add(parts)
	for part:=0;part<parts;part++{
		go wm.part(url, part, wg, err)
	}
	wg.Wait()
}

func (wm *workerMultipart) part(url string, part int, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if bytes,be:=wm.upload.bytes(wm.self, part);nil!=be{
		*err=be
	} else {
		wm.params.http.R().SetBody(bytes).Put(url)
	}
}
