package service

import (
	"example/modules/example/domain/entity"
	"example/modules/example/domain/repository"
)

type ExampleService struct {
	// тут зависимости
	repository repository.ExampleRepository
}

func Create(repository repository.ExampleRepository) *ExampleService {
	return &ExampleService{
		repository: repository,
	}
}

func (adapter *ExampleService) Push(name entity.Name) (string, error) {
	// тут бизнес логика
	result, err := adapter.repository.Push(name)
	if err != nil {
		// ещё тут нужно логировать но щас лень было добавлять ещё один интерфейс и прокирывать его
		return "", err
	}
	return result, nil
}
