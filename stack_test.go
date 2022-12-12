package collections

import "testing"

func TestSliceStackPeek(t *testing.T) {
	stack := new(SliceStack[int])
	stack.Push(15)
	got, _ := stack.Peek()
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSliceStackPeekEmpty(t *testing.T) {
	stack := new(SliceStack[int])
	_, got := stack.Peek()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPushLifo(t *testing.T) {
	stack := new(SliceStack[int])
	stack.Push(10)
	stack.Push(20)
	got, _ := stack.Peek()
	want := 20

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPop(t *testing.T) {
	stack := new(SliceStack[int])
	stack.Push(10)
	stack.Push(20)
	got, _ := stack.Peek()
	want := 20

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	stack.Pop()
	got, _ = stack.Peek()
	want = 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPopEmpty(t *testing.T) {
	stack := new(SliceStack[int])
	_, got := stack.Peek()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEmptySliceStack(t *testing.T) {
	stack := new(SliceStack[int])

	got, _ := stack.Peek()
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = stack.Size()
	want = 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestIsEmpty(t *testing.T) {
	stack := new(SliceStack[int])

	got := stack.IsEmpty()
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsNotEmpty(t *testing.T) {
	stack := new(SliceStack[int])

	stack.Push(10)
	got := stack.IsEmpty()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestNoSize(t *testing.T) {
	stack := new(SliceStack[int])

	got := stack.Size()
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSize(t *testing.T) {
	stack := new(SliceStack[int])

	stack.Push(10)
	got := stack.Size()
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestElements(t *testing.T) {
	stack := new(SliceStack[int])

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	got := stack.Elements()
	want := []int{10, 20, 30}

	for k, v := range got {
		v2 := want[k]
		if v != v2 {
			t.Errorf("got %d, wanted %d", v, v2)
		}
	}
}
