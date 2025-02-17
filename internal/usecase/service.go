package usecase

import "project/internal/domain/entities"

// Exemplo de serviço de caso de uso
type ExampleService struct {
    // Dependências necessárias
}

func (s *ExampleService) Execute(entity entities.ExampleEntity) error {
    // Lógica do caso de uso
    return nil
} 