package main

type CircularList[T any] struct {
	Data []T
	Size uint
}

func NewCircularList[T any](size uint) *CircularList[T] {
	return &CircularList[T]{
		Data: make([]T, size),
		Size: size,
	}
}

func (cl CircularList[T]) normalizeIndex(index int) uint {
	// todo: add support to get the quotient and remainder to know which octave things belong to
	if index > int(cl.Size) {
		return uint(index % int(cl.Size))
	} else {
		return uint(index)
	}
}

func (cl CircularList[T]) get(index uint) T {
	return cl.Data[cl.normalizeIndex(int(index))]
}

func (cl CircularList[T]) set(index uint, value T) {
	cl.Data[cl.normalizeIndex(int(index))] = value
}
