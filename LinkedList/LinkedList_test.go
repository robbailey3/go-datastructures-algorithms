package LinkedList

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.Append(tt.args.data)
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			if got := l.Contains(tt.args.data); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data      int
		afterData int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.InsertAfter(tt.args.data, tt.args.afterData)
		})
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data       int
		beforeData int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.InsertBefore(tt.args.data, tt.args.beforeData)
		})
	}
}

func TestLinkedList_PeekFirst(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			if got := l.PeekFirst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeekFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PeekLast(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			if got := l.PeekLast(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeekLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.Prepend(tt.args.data)
		})
	}
}

func TestLinkedList_Print(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.Print()
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.Remove(tt.args.data)
		})
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.RemoveFirst()
		})
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.RemoveLast()
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			l.Reverse()
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	type fields struct {
		head *node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList[int]{
				head: tt.fields.head,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedList[int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
