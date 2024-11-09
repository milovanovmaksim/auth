package service

// Используется для преобразование одного значений в другое.
type Into[T any] interface {
	Into() T
}
