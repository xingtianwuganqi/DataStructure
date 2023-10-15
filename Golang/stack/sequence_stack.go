package stack

/*
1.顺序栈
*/
import "fmt"

type Stack struct {
	items []int
}

// 入栈
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

// 出栈
func (s *Stack) pop() int {
	if s.isEmpty() {
		panic("stack is empty")
	}
	last := len(s.items) - 1
	poped := s.items[last]
	s.items = s.items[:last]
	return poped
}

// 判空
func (s *Stack) isEmpty() bool {
	if len(s.items) > 0 {
		return false
	} else {
		return true
	}
}

func (s *Stack) PrintStack() {
	for i := 0; i < len(s.items); i++ {
		var item = s.items[i]
		fmt.Printf("index:%d,value:%d\n", i, item)
	}
}

// 匹配{}
func matchingString(str string) bool {
	var stack = &Stack{}
	for i := 0; i < len(str); i++ {
		var value = str[i]
		if value == '{' {
			stack.push(1)
		} else if value == '}' {
			stack.pop()
		}
	}

	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

func StackTest() {
	stack := Stack{}

	for i := 0; i <= 5; i++ {
		stack.push(i)
	}

	stack.PrintStack()

	for i := 0; i <= 5; i++ {
		stack.pop()
	}

	isEmpty := stack.isEmpty()
	fmt.Println(isEmpty)
}

func MatingStringTest() {
	str := "{hello,{world}"
	finish := matchingString(str)
	fmt.Println(finish)
}
