package slice

import "container/heap"

type PriorityQueue[T any] struct {
	impl heapImpl[T]
}

func NewPriorityQueue[T any](less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{impl: heapImpl[T]{less: less}}
}

func (pq *PriorityQueue[T]) Push(v T) {
	heap.Push(&pq.impl, v)
}

func (pq *PriorityQueue[T]) Pop() T {
	v := heap.Pop(&pq.impl)
	return v.(T)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.impl.Len()
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.Len() == 0
}

type heapImpl[T any] struct {
	items []T
	less  func(T, T) bool
}

func (w heapImpl[T]) Len() int { return len(w.items) }

func (w heapImpl[T]) Less(i, j int) bool { return w.less(w.items[i], w.items[j]) }

func (w heapImpl[T]) Swap(i, j int) { w.items[i], w.items[j] = w.items[j], w.items[i] }

func (w *heapImpl[T]) Pop() any {
	n := len(w.items)
	value := w.items[n-1]
	w.items = w.items[:n-1]

	return value
}

func (w *heapImpl[T]) Push(v any) {
	w.items = append(w.items, v.(T))
}
