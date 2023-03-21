package hud

type worker interface {
	do() (err error)
}
