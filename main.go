package main

import (
	linear "datastructures/linear"
	"fmt"
)

func main() {

	testFixedStack()

}

func testFixedStack() {

	fixedStack := linear.StackFixed[int](3)
	val, err := fixedStack.Pop()
	if err != nil {
		fmt.Println(err, val)
	}

	fixedStack.Push(1)
	fixedStack.Push(2)
	fixedStack.Push(3)
	err = fixedStack.Push(4)

	if err != nil {
		fmt.Println(err)
	}

	fixedStack.Push(6)
	fixedStack.Push(5)
	fixedStack.Push(7)

	fmt.Printf("fixedStack: %v\n", fixedStack.Container)
}

func testStack() {

	que := linear.Stack[int]{}

	que.Container = append(que.Container, 1, 2, 4, 5)

	fmt.Printf("%v\n", que.Container)
	que.Push(6)

	x, y := que.Pop()
	fmt.Printf("%v %v\n", x, y)

	fmt.Printf("%v\n", que.Container)
	fmt.Printf("after pop %v\n", que.Container)

	que.Clear()
	fmt.Printf("ceared: %v\n", que.Container)

	que.Push(7)
	que.Push(8)
	que.Push(9)
	que.Push(10)

	fmt.Printf("peek: %v\n", que.Peek())
	fmt.Printf("first: %v\n", que.First())

	words := linear.Stack[string]{}

	words.Push("a")
	words.Push("b")
	words.Push("c")
	words.Push("e")

	fmt.Printf("words: %v\n", words)
}
