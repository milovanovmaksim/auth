package service

type To[T any] interface {
	To() T
}
