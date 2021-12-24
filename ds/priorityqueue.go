package ds

import (
	"container/heap"
	"fmt"
)

type PriorityQueue[T comparable] struct {
	pq *priorityQueue[T]
}

// NewPriorityQueue creates a new priority queue using the order given by the
// less function. The queue will return items in the ascending order, the
// lowest item will be returned first.
func NewPriorityQueue[T comparable](less func(a, b T) bool) *PriorityQueue[T] {
	pq := &priorityQueue[T]{
		less:  less,
		index: make(map[T]int),
	}
	heap.Init(pq)
	return &PriorityQueue[T]{pq: pq}
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pq.Len()
}

func (pq *PriorityQueue[T]) Push(x T) {
	heap.Push(pq.pq, x)
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(pq.pq).(T)
}

func (pq *PriorityQueue[T]) Update(x T) error {
	if i, ok := pq.pq.index[x]; ok {
		heap.Fix(pq.pq, i)
		return nil
	}
	return fmt.Errorf("priorityqueue: %v not present in priority queue", x)
}

type priorityQueue[T comparable] struct {
	heap  []T
	less  func(a, b T) bool
	index map[T]int
	nil   T
}

func (pq *priorityQueue[T]) Len() int { return len(pq.heap) }

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.heap[i], pq.heap[j])
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
	pq.index[pq.heap[i]], pq.index[pq.heap[j]] = i, j
}

func (pq *priorityQueue[T]) Push(x any) {
	n := len(pq.heap)
	pq.index[x.(T)] = n
	pq.heap = append(pq.heap, x.(T))
}

func (pq *priorityQueue[T]) Pop() any {
	old := pq.heap
	n := len(old)
	item := old[n-1]
	old[n-1] = pq.nil // avoid memory leak
	delete(pq.index, item)
	pq.heap = old[0 : n-1]
	return item
}
