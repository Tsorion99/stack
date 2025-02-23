package stack

import (
	"errors"
)

type stack[T any] struct {
	data []T
}

func MakeStack[T any](size uint) stack[T] {
	return stack[T]{data: make([]T, size)}
}

func (s *stack[T]) Push(elems ...T) {
	s.data = append(s.data, elems...)
}

func (s *stack[T]) Pop() error {
	if len(s.data) <= 0 {
		return errors.New("stack is empty")
	}
	s.data = s.data[:len(s.data)-1]
	return nil
}

func (s *stack[T]) Top() (*T, error) {
	if len(s.data) <= 0 {
		return nil, errors.New("stack is empty")
	}
	return &s.data[len(s.data)-1], nil
}

func (s *stack[T]) Len() int {
	return len(s.data)
}

func (s *stack[T]) Empty() bool {
	return len(s.data) == 0
}
