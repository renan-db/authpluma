package usecase

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"project/internal/modules/user/entity"
	interfaces "project/internal/modules/user/interface"
)

// Garantir que UserUseCase implementa UserInterface
var _ interfaces.UserInterface = (*UserUseCase)(nil)

// UserUseCase implementa a interface UserInterface
type UserUseCase struct {
	repo interfaces.UserInterface
}

// CRUD
func (u *UserUseCase) CreateUser(ctx context.Context, user *entity.UserEntity) error {
	if err := u.ValidateAll(user); err != nil {
		return err
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) GetUserById(ctx context.Context, id int32) (*entity.UserEntity, error) {
	return u.repo.GetUserById(ctx, id)
}

func (u *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *entity.UserEntity) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id int32) error {
	return u.repo.DeleteUser(ctx, id)
}

func (u *UserUseCase) ListUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return u.repo.ListUsers(ctx)
}

// Hooks
func (u *UserUseCase) BeforeCreate() time.Time {
	return time.Now()
}

func (u *UserUseCase) BeforeUpdate() time.Time {	
	return time.Now()
}

// Validations
 func (u *UserUseCase) ValidateRequiredFields(user *entity.UserEntity) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("nome é obrigatório")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email é obrigatório")
	}
	return nil
}

func (u *UserUseCase) ValidateNameLength(name string) error {
	length := len(strings.TrimSpace(name))
	if length < 3 {
		return errors.New("nome deve ter no mínimo 3 caracteres")
	}
	if length > 256 {
		return errors.New("nome deve ter no máximo 256 caracteres")
	}
	return nil
}

func (u *UserUseCase) ValidateNameFormat(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("nome não pode estar vazio")
	}
	return nil
}

func (u *UserUseCase) ValidateEmailLength(email string) error {
	length := len(strings.TrimSpace(email))
	if length < 5 { // exemplo@x.y
		return errors.New("email deve ter no mínimo 5 caracteres")
	}
	if length > 256 {
		return errors.New("email deve ter no máximo 256 caracteres")
	}
	return nil
}

func (u *UserUseCase) ValidateEmailFormat(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("formato de email inválido")
	}
	return nil
}

func (u *UserUseCase) ValidateDatabaseConnection() error {
	// Implementar verificação de conexão
	return nil
}

func (u *UserUseCase) ValidateEmailExistsInDB(email string) error {
	// Implementar verificação de email existente
	user, err := u.GetUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("email já cadastrado")
	}
	return nil
}

func (u *UserUseCase) ValidateAll(user *entity.UserEntity) error {
	if err := u.ValidateRequiredFields(user); err != nil {
		return err
	}
	if err := u.ValidateNameLength(user.Name); err != nil {
		return err
	}
	if err := u.ValidateNameFormat(user.Name); err != nil {
		return err
	}
	if err := u.ValidateEmailLength(user.Email); err != nil {
		return err
	}
	if err := u.ValidateEmailFormat(user.Email); err != nil {
		return err
	}
	if err := u.ValidateDatabaseConnection(); err != nil {
		return err
	}
	if err := u.ValidateEmailExistsInDB(user.Email); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) UserValidateAll(user *entity.UserEntity) error {
	// Chama o método de validação existente
	return u.ValidateAll(user)
	
}

func (u *UserUseCase) ValidateUserCreation(user *entity.UserEntity) (*entity.UserEntity, error) {
	// Realiza a validação completa
	if err := u.ValidateAll(user); err != nil {
		return nil, err // Retorna nil e o erro se a validação falhar
	}
	// Retorna o usuário e nenhum erro se a validação passar
	return user, nil
}

