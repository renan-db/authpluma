package interfaces

type ValidateStruct interface {
	Validate(req interface{}) error
}
