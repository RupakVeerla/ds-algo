package doublelinkedlist

type node struct {
	value    interface{}
	next     *node
	previous *node
}

type doublelinkedlist struct {
	head   *node
	tail   *node
	length int
}

func New(val interface{}) *doublelinkedlist {
	n := &node{
		value:    val,
		next:     nil,
		previous: nil,
	}
	l := &doublelinkedlist{
		head:   n,
		tail:   n,
		length: 1,
	}
	return l
}

func (l *doublelinkedlist) Append(val interface{}) {
	n := &node{
		value:    val,
		next:     nil,
		previous: l.tail,
	}
	l.tail.next = n
	l.tail = n
	l.length++
	return
}

func (l *doublelinkedlist) Prepend(val interface{}) {
	n := &node{
		value:    val,
		next:     l.head,
		previous: nil,
	}
	l.head = n
	l.length++
}

func (l *doublelinkedlist) Insert(val interface{}, index int) {
	if index == 0 {
		l.Prepend(val)
		return
	}
	if index >= l.length {
		l.Append(val)
		return
	}
	var n *node
	if l.length/2 > index {
		n = l.head
		for i := 1; i < index; i++ {
			n = n.next
		}
	} else {
		n = l.tail
		for i := 1; i < l.length-index; i++ {
			n = n.previous
		}
	}
	next := n.next
	previous := n.previous
	node := &node{val, next, previous}
	n.previous = node
	n.next = node
	l.length++
}

func (l *doublelinkedlist) Delete(index int) {
	if index == 0 {
		l.head = l.head.next
		l.head.previous = nil
		l.length--
		return
	}
	if index == l.length-1 {
		l.tail = l.tail.previous
		l.tail.next = nil
		l.length--
		return
	}
	if index > l.length-1 {
		return
	}
	var n *node
	if l.length/2 > index {
		n := l.head
		for i := 1; i < index; i++ {
			n = n.next
		}
	} else {
		n := l.tail
		for i := 1; i < l.length-index; i++ {
			n = n.previous
		}
	}
	n.next.next.previous = n
	n.next = n.next.next
	l.length--
}

func (l *doublelinkedlist) PrintList() (lst []interface{}) {
	h := l.head
	for h != nil {
		lst = append(lst, h.value)
		h = h.next
	}
	return
}
