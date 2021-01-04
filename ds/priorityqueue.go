package ds

import (
	"container/heap"

	"github.com/go-errors/errors"
)

type priorityQueue struct {
	heap     []interface{}
	priority func(interface{}) int
	index    map[interface{}]int
}

type PriorityQueue struct {
	pq *priorityQueue
}

// NewPriorityQueue creates a new priority queue for the given priority
// function. Items will be sorted in the ascending order, an item with the
// lowest priority will be poped first.
func NewPriorityQueue(priority func(interface{}) int) *PriorityQueue {
	pq := &priorityQueue{
		priority: priority,
		index:    make(map[interface{}]int),
	}
	heap.Init(pq)
	return &PriorityQueue{pq: pq}
}

func (pq *PriorityQueue) Len() int {
	return pq.pq.Len()
}

func (pq *PriorityQueue) Push(x interface{}) {
	heap.Push(pq.pq, x)
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(pq.pq)
}

func (pq *PriorityQueue) Update(x interface{}) error {
	if i, ok := pq.pq.index[x]; ok {
		heap.Fix(pq.pq, i)
		return nil
	}
	return errors.Errorf("%v not present in priority queue", x)
}

func (pq *priorityQueue) Len() int { return len(pq.heap) }

func (pq *priorityQueue) Less(i, j int) bool {
	return pq.priority(pq.heap[i]) < pq.priority(pq.heap[j])
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
	pq.index[pq.heap[i]] = i
	pq.index[pq.heap[j]] = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(pq.heap)
	pq.index[x] = n
	pq.heap = append(pq.heap, x)
}

func (pq *priorityQueue) Pop() interface{} {
	old := pq.heap
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	delete(pq.index, item)
	pq.heap = old[0 : n-1]
	return item
}
