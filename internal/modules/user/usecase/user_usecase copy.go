package usecase

// import (
// 	"context"
// 	"errors"
// 	"regexp"
// 	"strings"
// 	"time"
// )

// // Garantir que UserUseCase implementa UserInterface
// var _ UserUseCase = (*UserInterface)(nil)

// // UserUseCase implementa a interface UserInterface
// type UserUseCase struct {
// 	// Adicione os campos necessários, como um repositório
// 	repo UserRepository // Exemplo de repositório
// }

// // CRUD
// func (u *UserUseCase) Create(ctx context.Context, user *UserEntity) error {
// 	if err := u.ValidateAll(user); err != nil {
// 		return err
// 	}
// 	return u.repo.Create(ctx, user)
// }

// func (u *UserUseCase) GetByID(ctx context.Context, id int32) (*UserEntity, error) {
// 	return u.repo.GetByID(ctx, id)
// }

// func (u *UserUseCase) Update(ctx context.Context, user *UserEntity) error {
// 	return u.repo.Update(ctx, user)
// }

// func (u *UserUseCase) Delete(ctx context.Context, id int32) error {
// 	return u.repo.Delete(ctx, id)
// }

// func (u *UserUseCase) List(ctx context.Context) ([]UserEntity, error) {
// 	return u.repo.List(ctx)
// }

// // Hooks
// func (u *UserUseCase) BeforeCreate() time.Time {
// 	return time.Now()
// }

// func (u *UserUseCase) BeforeUpdate() time.Time {
// 	return time.Now()
// }

// // Validações
// func (u *UserUseCase) ValidateRequiredFields(user *UserEntity) error {
// 	if strings.TrimSpace(user.Name) == "" {
// 		return errors.New("nome é obrigatório")
// 	}
// 	if strings.TrimSpace(user.Email) == "" {
// 		return errors.New("email é obrigatório")
// 	}
// 	return nil
// }

// func (u *UserUseCase) ValidateNameLength(name string) error {
// 	length := len(strings.TrimSpace(name))
// 	if length < 3 {
// 		return errors.New("nome deve ter no mínimo 3 caracteres")
// 	}
// 	if length > 256 {
// 		return errors.New("nome deve ter no máximo 256 caracteres")
// 	}
// 	return nil
// }

// func (u *UserUseCase) ValidateNameFormat(name string) error {
// 	if strings.TrimSpace(name) == "" {
// 		return errors.New("nome não pode estar vazio")
// 	}
// 	return nil
// }

// func (u *UserUseCase) ValidateEmailLength(email string) error {
// 	length := len(strings.TrimSpace(email))
// 	if length < 5 { // exemplo@x.y
// 		return errors.New("email deve ter no mínimo 5 caracteres")
// 	}
// 	if length > 256 {
// 		return errors.New("email deve ter no máximo 256 caracteres")
// 	}
// 	return nil
// }

// func (u *UserUseCase) ValidateEmailFormat(email string) error {
// 	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	if !emailRegex.MatchString(email) {
// 		return errors.New("formato de email inválido")
// 	}
// 	return nil
// }

// func (u *UserUseCase) ValidateDatabaseConnection() error {
// 	// Implementar verificação de conexão
// 	return nil
// }

// func (u *UserUseCase) ValidateEmailExistsInDB(email string) error {
// 	// Implementar verificação de email existente
// 	return u.repo.CheckEmailExists(email)
// }

// func (u *UserUseCase) ValidateAll(user *UserEntity) error {
// 	if err := u.ValidateRequiredFields(user); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateNameLength(user.Name); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateNameFormat(user.Name); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateEmailLength(user.Email); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateEmailFormat(user.Email); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateDatabaseConnection(); err != nil {
// 		return err
// 	}
// 	if err := u.ValidateEmailExistsInDB(user.Email); err != nil {
// 		return err
// 	}
// 	return nil
// }