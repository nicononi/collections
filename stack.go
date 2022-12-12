package collections

// Collection that can accumulate elements of type T
// respecting the Last-In-First-Out pattern
type SliceStack[T comparable] struct {
	elements []T
	head     int
}

// Pushes the given element T at the top of the Stack
func (stack *SliceStack[T]) Push(element T) {
	stack.head++
	stack.elements = append(stack.elements, element)
}

// Removes the element T at the peek of the Stack
// and returns it back to the client
func (stack *SliceStack[T]) Pop() (T, bool) {
	if stack.head == 0 {
		var result T
		return result, false
	}

	element := stack.elements[stack.head-1]
	stack.head--
	return element, true
}

// Returns the element T sitting at the peek of the Stack
func (stack *SliceStack[T]) Peek() (T, bool) {
	if stack.head == 0 {
		var result T
		return result, false
	}
	return stack.elements[stack.head-1], true
}

// Returns the Slice containing the elements within the Stack
func (stack *SliceStack[T]) Elements() []T {
	return stack.elements
}

// Returns the amount of elements within the Stack
func (stack *SliceStack[T]) Size() int {
	return stack.head
}

// Returns the amount of elements within the Stack
func (stack *SliceStack[T]) IsEmpty() bool {
	return stack.Size() == 0
}
