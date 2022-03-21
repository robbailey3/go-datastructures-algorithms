package LinkedList

import "fmt"

type nodeValue interface {
	// Just using these because they feel sensible
	int | string | float64
}

type node[T nodeValue] struct {
	data T
	next *node[T]
}

type LinkedList[T nodeValue] struct {
	head *node[T]
}

func New[T nodeValue]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(data T) {
	n := &node[T]{data, nil}
	if l.head == nil {
		l.head = n
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = n
	}
}

func (l *LinkedList[T]) Prepend(data T) {
	n := &node[T]{data, l.head}
	l.head = n
}

func (l *LinkedList[T]) Remove(data T) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	cur := l.head
	for cur.next != nil {
		if cur.next.data == data {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

func (l *LinkedList[T]) RemoveLast() {
	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		return
	}
	cur := l.head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
}

func (l *LinkedList[T]) Contains(data T) bool {
	if l.head == nil {
		return false
	}
	cur := l.head
	for cur != nil {
		if cur.data == data {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList[T]) InsertAfter(data T, afterData T) {
	if l.head == nil {
		return
	}
	cur := l.head
	for cur != nil {
		if cur.data == afterData {
			n := &node[T]{data, cur.next}
			cur.next = n
			return
		}
		cur = cur.next
	}
}

func (l *LinkedList[T]) InsertBefore(data T, beforeData T) {
	if l.head == nil {
		return
	}
	if l.head.data == beforeData {
		n := &node[T]{data, l.head}
		l.head = n
		return
	}
	cur := l.head
	for cur.next != nil {
		if cur.next.data == beforeData {
			n := &node[T]{data, cur.next}
			cur.next = n
			return
		}
		cur = cur.next
	}
}

func (l *LinkedList[T]) Reverse() {
	if l.head == nil {
		return
	}
	var prev *node[T]
	cur := l.head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

func (l *LinkedList[T]) Print() {
	if l.head == nil {
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Print(cur.data)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) PeekFirst() T {
	if l.head == nil {
		return *new(T)
	}
	return l.head.data
}

func (l *LinkedList[T]) PeekLast() T {
	if l.head == nil {
		return *new(T)
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur.data
}

func (l *LinkedList[T]) Size() int {
	size := 0
	cur := l.head
	for cur != nil {
		size++
		cur = cur.next
	}
	return size
}
