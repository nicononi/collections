package collections

import "testing"

func TestQueuePeek(t *testing.T) {
	queue := new(SliceQueue[int])
	queue.Enqueue(15)
	got, _ := queue.Peek()
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	queue := new(SliceQueue[int])
	_, got := queue.Peek()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEnqueueFifo(t *testing.T) {
	queue := new(SliceQueue[int])
	queue.Enqueue(10)
	queue.Enqueue(20)
	got, _ := queue.Peek()
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDequeue(t *testing.T) {
	queue := new(SliceQueue[int])
	queue.Enqueue(10)
	queue.Enqueue(20)
	got, _ := queue.Peek()
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	queue.Dequeue()
	got, _ = queue.Peek()
	want = 20

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestDequeueEmpty(t *testing.T) {
	queue := new(SliceQueue[int])
	_, got := queue.Dequeue()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEmptyQueue(t *testing.T) {
	queue := new(SliceQueue[int])

	got, _ := queue.Peek()
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = queue.Size()
	want = 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
