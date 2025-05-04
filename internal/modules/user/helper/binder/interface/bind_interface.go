// Package interfaces define as interfaces do submódulo de bind do módulo de usuário.
package interfaces

// Binder representa a interface para realizar o bind de dados em uma estrutura.
type Binder interface {
	Bind(target interface{}) error
}
