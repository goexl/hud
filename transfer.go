package hud

// Transfer 传输器
type Transfer struct {
	params *params
}

func newTransfer() *Transfer {
	return &Transfer{
		params: newParams(),
	}
}

func (t *Transfer) Upload() *uploadBuilder {
	return newUploadBuilder(t.params)
}
