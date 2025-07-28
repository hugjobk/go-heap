package heap

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type Lessable[T any] interface {
	Less(T) bool
}

type HeapInterface interface {
	heap.Interface
	Top() any
}

type OrderedHeap[T any] []T

func (h OrderedHeap[T]) Len() int {
	return len(h)
}

func (h OrderedHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *OrderedHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *OrderedHeap[T]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h OrderedHeap[T]) Top() any {
	return h[0]
}

type MinHeap[T constraints.Ordered] struct {
	OrderedHeap[T]
}

func (h MinHeap[T]) Less(i, j int) bool {
	return h.OrderedHeap[i] < h.OrderedHeap[j]
}

type MaxHeap[T constraints.Ordered] struct {
	OrderedHeap[T]
}

func (h MaxHeap[T]) Less(i, j int) bool {
	return h.OrderedHeap[i] > h.OrderedHeap[j]
}

type LessHeap[T Lessable[T]] struct {
	OrderedHeap[T]
}

func (h LessHeap[T]) Less(i, j int) bool {
	return h.OrderedHeap[i].Less(h.OrderedHeap[j])
}

type Heap[T any] struct {
	h HeapInterface
}

func (h Heap[T]) Len() int {
	return h.h.Len()
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.h, x)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

func (h Heap[T]) Top() T {
	return h.h.Top().(T)
}

func NewMinHeap[T constraints.Ordered](data ...T) *Heap[T] {
	h := &MinHeap[T]{data}
	heap.Init(h)
	return &Heap[T]{h}
}

func NewMaxHeap[T constraints.Ordered](data ...T) *Heap[T] {
	h := &MaxHeap[T]{data}
	heap.Init(h)
	return &Heap[T]{h}
}

func NewLessHeap[T Lessable[T]](data ...T) *Heap[T] {
	h := &LessHeap[T]{data}
	heap.Init(h)
	return &Heap[T]{h}
}
