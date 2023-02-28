package stack

import "fmt"

type item struct {
	value string // hold any data type
	next  *item
}
type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Push(v string) {
	stack.top = &item{
		value: v,
		next:  stack.top,
	}
	stack.size++
}
func (stack *Stack) Len() int {
	return stack.size
}
func (stack *Stack) isEmpty() bool {
	return stack.Len() == 0
}
func (stack *Stack) Pop() string {
	if stack.isEmpty() {
		return ""
	}

	valueToPop := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return valueToPop
}
func (stack *Stack) Peek() string {
	if stack.isEmpty() {
		return ""
	}
	return stack.top.value
}

func (stack *Stack) ReadStack() {
	temp := stack.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println(" ")
}
