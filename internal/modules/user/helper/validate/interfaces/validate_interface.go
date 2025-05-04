// Package interfaces define as interfaces do submódulo de validate do módulo de usuário.
package interfaces

// Validater define a interface para validar uma estrutura.
type Validater interface {
	Validate(data interface{}) error
}
