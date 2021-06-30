package queues

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type queue struct {
	first  *node
	last   *node
	length int
}

func New() *queue {
	return &queue{nil, nil, 0}
}

func (q *queue) Peek() interface{} {
	return q.first.value
}

func (q *queue) Enqueue(val interface{}) {
	n := &node{val, nil}
	if q.first == nil {
		q.first = n
		q.last = n
		q.length++
		return
	}
	q.last.next = n
	q.last = n
	q.length++
}

func (q *queue) Dequeue() (val interface{}) {
	val = q.first.value
	q.first = q.first.next
	q.length--
	return
}

func (q *queue) Print() {
	var l []interface{}
	n := q.first
	for i := 0; i < q.length; i++ {
		l = append(l, n.value)
		n = n.next
	}
	fmt.Println(l)
}
