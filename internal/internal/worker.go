package internal

type Worker interface {
	Do() (err error)
}
