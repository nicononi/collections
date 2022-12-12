package collections

// Collection that can accumulate elements of type T
// respecting the First-In-First-Out pattern
type SliceQueue[T comparable] struct {
	elements []T
}

// Enqueues the given element T at the bottom of the Queue
func (queue *SliceQueue[T]) Enqueue(element T) {
	queue.elements = append(queue.elements, element)
}

// Removes the element T at the peek of the Queue
// and returns it back to the client
func (queue *SliceQueue[T]) Dequeue() (T, bool) {
	if queue.Size() == 0 {
		var result T
		return result, false
	}

	element := queue.elements[0]
	queue.elements = queue.elements[1:]
	return element, true
}

// Returns the element T sitting at the peek of the Queue
func (queue *SliceQueue[T]) Peek() (T, bool) {
	if queue.Size() == 0 {
		var result T
		return result, false
	}

	return queue.elements[0], true
}

// Returns the Slice containing the elements within the Queue
func (queue *SliceQueue[T]) Elements() []T {
	return queue.elements
}

// Returns the amount of elements within the Queue
func (queue *SliceQueue[T]) Size() int {
	return len(queue.elements)
}

// Returns the amount of elements within the Queue
func (queue *SliceQueue[T]) IsEmpty() bool {
	return queue.Size() == 0
}
