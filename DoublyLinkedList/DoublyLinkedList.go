package doublylinkedlist

type nodeValue interface {
	// Just using these because they feel sensible
	int | string | float64
}

type node[T nodeValue] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type DoublyLinkedList[T nodeValue] struct {
	head *node[T]
	tail *node[T]
}

func New[T nodeValue](value T) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList[T]) Append(newNode *node[T]) {
	if d.head == nil {
		d.head = newNode
		d.tail = d.head
	} else {
		d.tail.next = newNode
		d.tail = d.tail.next
	}
}

func (d *DoublyLinkedList[T]) Prepend(newNode *node[T]) {
	if d.head == nil {
		d.head = newNode
		d.tail = d.head
	} else {
		d.head = newNode
		d.head.next.prev = d.head
	}
}

func (d *DoublyLinkedList[T]) Remove(node *node[T]) {
	if d.head == nil {
		return
	}
	if d.head == node {
		d.head = d.head.next
		return
	}
	cur := d.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (d *DoublyLinkedList[T]) RemoveFirst() {
	if d.head == nil {
		return
	}
	d.head = d.head.next
}

func (d *DoublyLinkedList[T]) RemoveLast() {
	if d.head == nil {
		return
	}
	d.tail = d.tail.prev
}

func (d *DoublyLinkedList[T]) GetHead() *node[T] {
	return d.head
}

func (d *DoublyLinkedList[T]) GetTail() *node[T] {
	return d.tail
}

func (d *DoublyLinkedList[T]) Size() int {
	if d.head == nil {
		return 0
	}
	size := 1
	cur := d.head
	for cur.next != nil {
		size++
		cur = cur.next
	}
	return size
}
