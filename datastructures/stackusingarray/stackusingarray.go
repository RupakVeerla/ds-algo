package stackusingarray

type item interface{}

type Stack struct {
	value []item
}

func New() (s *Stack) {
	s = &Stack{value: []item{}}
	return
}

func (s *Stack) Push(val interface{}) {
	s.value = append(s.value, val)
}

func (s *Stack) Peek() interface{} {
	return s.value[len(s.value)-1]
}

func (s *Stack) Pop() (val interface{}) {
	val = s.value[len(s.value)-1]
	s.value = append(s.value[:len(s.value)-1])
	return
}
