package service

import "fmt"

type LengthPasswordError struct {
	length int64
}

func NewLengthPasswordError(length int64) LengthPasswordError {
	return LengthPasswordError{length: length}
}

func (e LengthPasswordError) Error() string {
	return fmt.Sprintf("password must be more then %v characters", e.length)
}

type EmptyPasswordError struct{}

func NewEmptyPasswordError() EmptyPasswordError {
	return EmptyPasswordError{}
}

func (e EmptyPasswordError) Error() string {
	return "password is empty"
}

type ConfirmationPasswordError struct{}

func NewConfirmationPasswordError() ConfirmationPasswordError {
	return ConfirmationPasswordError{}
}

func (e ConfirmationPasswordError) Error() string {
	return "password and password_confirm must be the same"
}
