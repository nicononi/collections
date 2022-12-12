package collections

// Collection that can accumulate elements of type T
// respecting the single list pattern
type SliceList[T comparable] struct {
	elements []T
}

// Appends the given element T at the end of the list
func (list *SliceList[T]) Append(element T) {
	list.elements = append(list.elements, element)
}

// Removes the element T at the given index within the list
// and returns it back to the client
func (list *SliceList[T]) Remove(index int) (T, bool) {
	if list.Size() == 0 {
		var result T
		return result, false
	}

	element := list.elements[index]
	newElements := make([]T, 0)
	newElements = append(newElements, list.elements[:index]...)
	list.elements = append(newElements, list.elements[index+1:]...)
	return element, true
}

// Returns the element T at the given index within the list
func (list *SliceList[T]) Get(index int) (T, bool) {
	if list.Size() == 0 {
		var result T
		return result, false
	}

	return list.elements[0], true
}

// Checks if the given element T exists within the list
func (list *SliceList[T]) Contains(element T) bool {
	if list.Size() == 0 {
		return false
	}

	for _, v := range list.elements {
		if v == element {
			return true
		}
	}

	return false
}

// Returns the Slice containing the elements within the list
func (list *SliceList[T]) Elements() []T {
	return list.elements
}

// Returns the amount of elements within the list
func (list *SliceList[T]) Size() int {
	return len(list.elements)
}

// Returns the amount of elements within the list
func (list *SliceList[T]) IsEmpty() bool {
	return list.Size() == 0
}
