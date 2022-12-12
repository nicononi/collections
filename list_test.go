package collections

import "testing"

func TestListPeek(t *testing.T) {
	list := new(SliceList[int])
	list.Append(15)
	got, _ := list.Get(0)
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestListPeekEmpty(t *testing.T) {
	list := new(SliceList[int])
	_, got := list.Get(0)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestRemove(t *testing.T) {
	list := new(SliceList[int])
	list.Append(10)
	list.Append(20)
	got, _ := list.Get(0)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got, _ = list.Remove(0)
	want = 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestContains(t *testing.T) {
	list := new(SliceList[int])

	list.Append(10)
	got := list.Contains(10)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestNotContains(t *testing.T) {
	list := new(SliceList[int])

	list.Append(20)
	got := list.Contains(10)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
