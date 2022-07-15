package util

type Mapper[T any] interface {
	mapTo() T
}
