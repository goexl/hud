package hud

type lifecycle interface {
	Request()
	Initiate() (id string)
	Abort()
	Complete()
}
