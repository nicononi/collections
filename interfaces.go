package collections

// Represents something that can append elements of any type
type Queue[T any] interface {
	Collectionable[T]
	Enqueue(element T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}

// Represents something that can append elements of any type
type Stack[T any] interface {
	Collectionable[T]
	Push(element T)
	Pop() (T, bool)
	Peek() (T, bool)
}

// Represents something that can append elements of any type
type List[T any] interface {
	Collectionable[T]
	Append(element T)
	Get(index int) (T, bool)
	Contains(element T) bool
}

// Represents something that has elements of any type
type Collectionable[T any] interface {
	Elements() []T
	Size() int
	IsEmpty() bool
}
