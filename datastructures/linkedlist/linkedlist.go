package linkedlist

type node struct {
	value interface{}
	next  *node
}

type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func New(val interface{}) *linkedlist {
	n := &node{
		value: val,
		next:  nil,
	}
	l := &linkedlist{
		head:   n,
		tail:   n,
		length: 1,
	}
	return l
}

func (l *linkedlist) Append(val interface{}) {
	n := &node{
		value: val,
		next:  nil,
	}
	l.tail.next = n
	l.tail = n
	l.length++
	return
}

func (l *linkedlist) Prepend(val interface{}) {
	n := &node{
		value: val,
		next:  l.head,
	}
	l.head = n
	l.length++
}

func (l *linkedlist) Insert(val interface{}, index int) {
	if index == 0 {
		l.Prepend(val)
		return
	}
	if index >= l.length-1 {
		l.Append(val)
		return
	}
	h := l.head
	for i := 1; i < index; i++ {
		h = h.next
	}
	temp := h.next
	n := &node{val, temp}
	h.next = n
	l.length++
}

func (l *linkedlist) Delete(index int) {
	if index == 0 {
		l.head = l.head.next
		l.length--
		return
	}
	if index > l.length-1 {
		return
	}
	h := l.head
	for i := 1; i < index; i++ {
		h = h.next
	}
	h.next = h.next.next
	l.length--
}

func (l *linkedlist) PrintList() (lst []interface{}) {
	h := l.head
	for h != nil {
		lst = append(lst, h.value)
		h = h.next
	}
	return
}

func (l *linkedlist) Reverse() {
	if l.length == 1 {
		return
	}
	h := l.head
	l.tail = l.head
	s := h.next
	for s != nil {
		t := s.next
		s.next = h
		h = s
		s = t
	}
	l.head.next = nil
	l.head = h
}
