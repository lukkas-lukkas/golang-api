package domain

type Command interface {
	Init([]string) error
	Exec() error
	Name() string
}
