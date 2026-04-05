package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	capasity  int
	container []T
}

func (s *Stack[T]) push(value T) {
	if s.capasity > 0 && (len(s.container) >= s.capasity) {
		panic("Error max capasity exeded")
	}

	s.container = append(s.container, value)
}

func (s *Stack[T]) pop() (T, error) {

	var zero T

	if len(s.container) <= 0 {
		return zero, errors.New("Empty Stack")
	}

	last := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]

	return last, nil
}

func (s *Stack[T]) clear() {
	s.container = []T{}
}

func (s *Stack[T]) peek() T {

	return s.container[len(s.container)-1]
}

func (s *Stack[T]) first() T {

	return s.container[0]
}

func StackFixed[T any](value int) *Stack[T] {

	return &Stack[T]{
		capasity:  value,
		container: make([]T, 0, value),
	}

}

func (s *Stack[T]) setSize(value int) {
	s.capasity = value
}

func main() {

	testFixedStack()

}

func testFixedStack() {

	fixedStack := StackFixed[int](3)

	fixedStack.push(1)
	fixedStack.push(2)
	fixedStack.push(3)
	fixedStack.pop()
	fixedStack.push(4)
	// fixedStack.push(5)
	// fixedStack.push(6)

	fmt.Printf("fixedStack: %v\n", fixedStack.container)
}

func testStack() {

	que := Stack[int]{}

	que.container = append(que.container, 1, 2, 4, 5)

	fmt.Printf("%v\n", que.container)
	que.push(6)

	x, y := que.pop()
	fmt.Printf("%v %v\n", x, y)

	fmt.Printf("%v\n", que.container)
	fmt.Printf("after pop %v\n", que.container)

	que.clear()
	fmt.Printf("ceared: %v\n", que.container)

	que.push(7)
	que.push(8)
	que.push(9)
	que.push(10)

	fmt.Printf("peek: %v\n", que.peek())
	fmt.Printf("first: %v\n", que.first())

	words := Stack[string]{}

	words.push("a")
	words.push("b")
	words.push("c")
	words.push("e")

	fmt.Printf("words: %v\n", words)
}
