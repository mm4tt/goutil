package ds

type Queue struct {
	data          []interface{}
	head, tail, n int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Len() int {
	return q.n
}

func (q *Queue) Enqueue(e interface{}) {
	if q.n == len(q.data) {
		q.resize()
	}
	q.data[q.tail] = e
	q.n++
	q.tail = (q.tail + 1) % len(q.data)
}

func (q *Queue) Dequeue() interface{} {
	if q.n == 0 {
		panic("queue is empty!")
	}
	e := q.data[q.head]
	q.data[q.head] = nil // avoid memory leak
	q.head = (q.head+1)%len(q.data)
	q.n--
	return e
}

func (q *Queue) SubQueue(n int) *Queue {
	if n > q.n {
		panic("queue is not big enough!")
	}
	q1 := &Queue{data: make([]interface{}, n), tail: 0, n: n}
	for i := 0; i < n; i++ {
		q1.data[i] = q.data[(i+q.head)%len(q.data)]
	}
	return q1
}

func (q *Queue) resize() {
	data := make([]interface{}, 2*len(q.data)+1)
	for i := 0; i < q.n; i++ {
		data[i] = q.data[(i+q.head)%len(q.data)]
	}
	q.data, q.head, q.tail = data, 0, len(q.data)
}

func (q Queue) Peek() interface{} {
	return q.PeekI(0)
}

// PeekI returns the ith element in the queue.
func (q Queue) PeekI(i int) interface{} {
	return q.data[(i+q.head) % len(q.data)]
}
