package LinkedList

import (
	"reflect"
	"testing"
)

func CreateLinkedList(length int) *LinkedList[int] {
	ll := New[int]()
	for i := 0; i < length; i++ {
		ll.Append(i)
	}
	return ll
}

func TestLinkedList_Append(t *testing.T) {
	t.Run("should increase size by one", func(t *testing.T) {
		ll := New[int]()
		ll.Append(1)
		if ll.Size() != 1 {
			t.Errorf("expected size to be 1, got %d", ll.Size())
		}
		ll.Append(2)
		if ll.Size() != 2 {
			t.Errorf("expected size to be 2, got %d", ll.Size())
		}
	})

	t.Run("should populate the head if not already defined", func(t *testing.T) {
		ll := New[int]()
		ll.Append(1)
		if ll.Size() != 1 {
			t.Errorf("expected size to be 1, got %d", ll.Size())
		}
		if ll.head.data != 1 {
			t.Errorf("expected head to be 1, got %d", ll.head.data)
		}
	})
}

func TestLinkedList_Contains(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name       string
		linkedList *LinkedList[int]
		args       args
		want       bool
	}{
		{
			name:       "should return true if the linked list contains the data",
			linkedList: CreateLinkedList(2),
			args: args{
				data: 1,
			},
			want: true,
		},
		{
			name:       "should return false if the linked list does not contain the data",
			linkedList: CreateLinkedList(2),
			args: args{
				data: 3,
			},
			want: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if got := testCase.linkedList.Contains(testCase.args.data); got != testCase.want {
				t.Errorf("Contains() = %v, want %v", got, testCase.want)
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
			if got := New[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
