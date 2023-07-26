package heap_test

import (
	"fmt"
	"testing"

	"github.com/hugjobk/go-heap"
)

// A implements Lessable interface so that it can be used with LessHeap
type A struct {
	i int
}

func (a A) Less(other A) bool {
	return a.i < other.i
}

func TestLessHeap(t *testing.T) {
	h := heap.NewLessHeap(A{1}, A{6}, A{3}, A{2}, A{5})
	h.Push(A{7})
	h.Push(A{2})
	h.Push(A{9})
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
}

func TestMinHeap(t *testing.T) {
	h := heap.NewMinHeap(1, 6, 3, 2, 5)
	h.Push(7)
	h.Push(2)
	h.Push(9)
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
}

func TestMaxHeap(t *testing.T) {
	h := heap.NewMaxHeap(1, 6, 3, 2, 5)
	h.Push(7)
	h.Push(2)
	h.Push(9)
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
}
