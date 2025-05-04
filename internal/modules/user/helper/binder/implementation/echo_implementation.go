// Package implementation define a implementação do submódulo de bind do módulo de usuário.
package implementation

import (
	"fmt"

	interfaces "project/internal/modules/user/helper/binder/interface"

	"github.com/labstack/echo/v4"
)

// Garante que echoBinder implementa a interface Binder.
var _ interfaces.Binder = (*echoBinder)(nil)

// echoBinder representa a implementação de bind utilizando Echo.
type echoBinder struct {
	ctx echo.Context
}

// NewEchoBinder cria uma nova instância de echoBinder com o contexto do Echo.
func NewEchoBinder(ctx echo.Context) *echoBinder {
	return &echoBinder{ctx: ctx}
}

// Bind realiza o bind dos dados da requisição para o destino informado.
func (e *echoBinder) Bind(target interface{}) error {
	if err := e.ctx.Bind(target); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}
	return nil
}
