package entity

import "errors"

type Name struct {
	value string
}

func Create(name string) (Name, error) {
	if len(name) == 0 {
		return Name{}, errors.New("Error value of class Name. Value: " + name)
	}
	return Name{value: name}, nil
}

func (name Name) GetValue() string {
	return name.value
}
