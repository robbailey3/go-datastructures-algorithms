package queue

type Queue[T any] struct {
	Items []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	return q.Items[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.Items)
}

func (q *Queue[T]) Clear() {
	q.Items = []T{}
}
