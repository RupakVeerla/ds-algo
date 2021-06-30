package stack

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type stack struct {
	top    *node
	bottom *node
	length int
}

func New(val interface{}) *stack {
	n := &node{val, nil}
	s := &stack{n, n, 1}
	return s
}

func (s *stack) Peek() interface{} {
	return s.top.value
}

func (s *stack) Push(val interface{}) {
	n := &node{val, s.top}
	s.top = n
	s.length++
}

func (s *stack) Pop() (v interface{}) {
	v = s.top.value
	s.top = s.top.next
	s.length--
	return
}

func (s *stack) Print() {
	stk := []interface{}{}
	x := s.top
	for i := 0; i < s.length; i++ {
		stk = append(stk, x.value)
		x = x.next
	}
	fmt.Println(stk)
}
