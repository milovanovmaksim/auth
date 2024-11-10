package service

// Используется для преобразования одного значения в другое.
type Into[T any] interface {
	Into() T
}
