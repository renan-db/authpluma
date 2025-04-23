package interfaces

type Bind interface {
	Binder(target interface{}) error
}
