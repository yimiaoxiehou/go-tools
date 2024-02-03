package queue

type Queue struct {
	items []interface{}
	limit int
}

// NewQueue creates a new queue with the specified size limit
func NewQueue(size int) Queue {
	return Queue{
		items: make([]interface{}, 0),
		limit: size,
	}
}

// Push adds an element to the queue, removing the first element if the queue is full.
func (q *Queue) Push(v interface{}) {
	if len(q.items) == q.limit {
		// Remove the first element if the queue is full.
		q.items = append(q.items[1:], v)
	}
	q.items = append(q.items, v)
}

// Pop removes and returns the first element from the queue
func (q *Queue) Pop() interface{} {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

// IsEmpty checks if the queue is empty and returns a boolean value.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// IsFull checks if the queue is full.
func (q *Queue) IsFull() bool {
	return len(q.items) == q.limit
}

// Has checks if the queue contains a specific element.
func (q *Queue) Has(v interface{}) bool {
	// Iterate over the items in the queue.
	for _, e := range q.items {
		// Check if the current item is equal to the specified element.
		if e == v {
			return true
		}
	}
	// Return false if the element is not found in the queue.
	return false
}
