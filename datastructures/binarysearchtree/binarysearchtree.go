package binarysearchtree

import (
	"fmt"
	"strconv"
)

type node struct {
	left  *node
	right *node
	value int
}

type binarysearchtree struct {
	root *node
}

func New() *binarysearchtree {
	return &binarysearchtree{}
}

func (b *binarysearchtree) Insert(val int) {
	n := &node{nil, nil, val}
	if b.root == nil {
		b.root = n
		return
	}
	cn := b.root
	for {
		if cn.value > val {
			if cn.left == nil {
				cn.left = n
				return
			}
			cn = cn.left
		} else {
			if cn.right == nil {
				cn.right = n
				return
			}
			cn = cn.right
		}
	}
}

func (b *binarysearchtree) Lookup(val int) bool {
	n := b.root
	for {
		if n.value == val {
			return true
		} else if n.value > val {
			if n.left == nil {
				return false
			}
			n = n.left
		} else {
			if n.right == nil {
				return false
			}
			n = n.right
		}
	}
}

func (b *binarysearchtree) Print() {
	b.root.Print()
}

func (n *node) Print() {
	fmt.Print(strconv.Itoa(n.value) + "   ")
	if n.left != nil {
		n.left.Print()
	}
	if n.right != nil {
		n.right.Print()
	}
}
