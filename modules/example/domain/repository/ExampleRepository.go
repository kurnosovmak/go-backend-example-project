package repository

import "example/modules/example/domain/entity"

type ExampleRepository interface {
	Push(name entity.Name) (string, error)
}
