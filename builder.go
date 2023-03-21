package hud

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}
