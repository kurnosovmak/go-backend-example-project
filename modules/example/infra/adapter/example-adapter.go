package adapter

import (
	"example/modules/example/domain/entity"
	"example/modules/example/domain/service"
	repository "example/modules/example/infra/reposiroty"
)

type ExampleAdapter struct {
	// тут зависимости
	service *service.ExampleService
}

func Create() *ExampleAdapter {
	return &ExampleAdapter{service: service.Create(repository.Create())}
}

func (adapter *ExampleAdapter) PushHelloWorld(nameRaw string) (string, error) {
	name, err := entity.Create(nameRaw)
	if err != nil {
		return "", err
	}
	return adapter.service.Push(name)
}
