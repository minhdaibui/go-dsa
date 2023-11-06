package linkedlist

import "fmt"

type node[T any] struct {
	Data T
	Next *node[T]
}

func newNode[T any](data T) *node[T] {
	return &node[T]{Data: data, Next: nil}
}

type Singly[T any] struct {
	Head *node[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (l *Singly[T]) InsertHead(data T) {
	nn := newNode[T](data)
	nn.Next = l.Head
	l.Head = nn
}

func (l *Singly[T]) InsertTail(data T) {
	if l.IsEmpty() {
		l.InsertHead(data)
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode[T](data)
}

func (l *Singly[T]) Transverse() {
	fmt.Print("[")
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Print(" ", cur.Data, " ")
	}
	fmt.Print("]")
	fmt.Println()
}

func (l *Singly[T]) IsEmpty() bool {
	return l.Head == nil
}
