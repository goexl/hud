package hud

// Transfer 传输器
type Transfer struct {
	params *params
}

func newTransfer(params *params) *Transfer {
	return &Transfer{
		params: params,
	}
}

func (t *Transfer) Upload() *uploadBuilder {
	return newUploadBuilder(t.params)
}
