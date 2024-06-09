package repository

import (
	"errors"
	"example/modules/example/domain/entity"
)

var repo *TextExampleRepository

type ExampleRepository interface {
	Push(name entity.Name) (string, error)
}

type TextExampleRepository struct {
	// тут зависимости
}

func Create() *TextExampleRepository {
	return &TextExampleRepository{
		// тут зависимости прокидываем их(их прокидываем из параметров метода)
	}
}

func (repo TextExampleRepository) Push(name entity.Name) (string, error) {
	// имитирую запрос каких-то данных
	if name.GetValue() == "error" {
		return "", errors.New("Error connection")
	}
	return "Hello world, " + name.GetValue(), nil
}
