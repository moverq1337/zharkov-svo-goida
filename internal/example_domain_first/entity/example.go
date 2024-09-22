package entity

import "errors"

type ExampleUser struct {
	Username string
	Age      int
}

func NewExampleUser(username string, age int) (*ExampleUser, error) {

	if username == "" {
		return nil, errors.New("username is empty")
	}

	if age <= 0 {
		return nil, errors.New("age is too small")
	}

	return &ExampleUser{
		Username: username,
		Age:      age,
	}, nil
}

func (e *ExampleUser) Rename(newName string) error {
	if newName == "" {
		return errors.New("newName is empty")
	}

	e.Username = newName
	return nil
}

func (e *ExampleUser) Birthday() error {
	e.Age++
	return nil
}
