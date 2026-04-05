package stack

import (
	"errors"
)

// Stack -> creates a Stack with no size
type Stack[T any] struct {
	Capacity  int
	Container []T
}

// StackFixed -> creates Stack with a fixed size
func StackFixed[T any](value int) *Stack[T] {

	return &Stack[T]{
		Capacity:  value,
		Container: make([]T, 0, value),
	}

}

// Push -> Appends value to stack returns error
// If using StackFixed errors when over capacity
func (s *Stack[T]) Push(value T) error {
	if s.Capacity > 0 && (len(s.Container) >= s.Capacity) {
		return errors.New("Error max capasity exeded")
	}

	s.Container = append(s.Container, value)
	return nil
}

// Pop -> returns value,error and removes last index
func (s *Stack[T]) Pop() (T, error) {

	var zero T

	if len(s.Container) <= 0 {
		return zero, errors.New("Error Empty Stack")
	}

	last := s.Container[len(s.Container)-1]
	s.Container = s.Container[:len(s.Container)-1]

	return last, nil
}

// Clear -> clears the stack
func (s *Stack[T]) Clear() {
	s.Container = []T{}
}

// Peek -> returns top item
func (s *Stack[T]) Peek() T {
	return s.Container[len(s.Container)-1]
}

// First -> returns bottom item
func (s *Stack[T]) First() T {
	return s.Container[0]
}
