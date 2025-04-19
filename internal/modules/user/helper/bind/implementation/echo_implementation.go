package implementation

import (
	"fmt"

	bindInterfaces "project/internal/modules/user/helper/bind/interface"

	"github.com/labstack/echo/v4"
)

// Garante que a struct implemente a interface esperada
var _ bindInterfaces.Bind = (*EchoBinder)(nil)

// Representa o binder com o echo
type EchoBinder struct {
	ctx echo.Context
}

// Cria uma nova instância com o contexto do echo
func NewEchoBinder(ctx echo.Context) *EchoBinder {
	return &EchoBinder{ctx: ctx}
}

// Implementa o binder com o echo
func (e *EchoBinder) Binder(target interface{}) error {
	if err := e.ctx.Bind(target); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}
	return nil
}
