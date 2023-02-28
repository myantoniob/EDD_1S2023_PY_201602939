package stack

type item struct {
	value interface{} // hold any data type
	next  *item
}
type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Push(v interface{}) {
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
func (stack *Stack) Pop() interface{} {
	if stack.isEmpty() {
		return nil
	}

	valueToPop := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return valueToPop
}
func (stack *Stack) Peek() interface{} {
	if stack.isEmpty() {
		return nil
	}
	return stack.top.value
}
