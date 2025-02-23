package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := MakeStack[int](0)
	want := 9
	s.Push(want)
	v, err := s.Top()

	if v == nil {
		t.Fatal("Top() return nil")
	}
	if *v != want || err != nil {
		t.Fatalf("Top() = %d, %v, want match for %d, nil", *v, err, want)
	}
}

func TestMultiPush(t *testing.T) {
	s := MakeStack[int](0)
	want := []int{9, 6, 3}
	s.Push(want[0], want[1], want[2])

	v, err := s.Top()
	if v == nil {
		t.Fatal("Top() return nil")
	}
	if *v != want[2] || err != nil {
		t.Fatalf("Top() = %d, %v, want match for %d, nil", *v, err, want[2])
	}

	s.Pop()
	v, err = s.Top()
	if v == nil {
		t.Fatal("Top() return nil")
	}
	if *v != want[1] || err != nil {
		t.Fatalf("Top() = %d, %v, want match for %d, nil", *v, err, want[1])
	}

	s.Pop()
	v, err = s.Top()
	if v == nil {
		t.Fatal("Top() return nil")
	}
	if *v != want[0] || err != nil {
		t.Fatalf("Top() = %d, %v, want match for %d, nil", *v, err, want[0])
	}
}

func TestTopNewStack(t *testing.T) {
	s := MakeStack[int](0)
	v, err := s.Top()

	if v != nil {
		t.Fatalf("Top() return not nil")
	}
	if err == nil || err.Error() != "stack is empty" {
		t.Fatalf("Top() = nil, nil, want match for nil, \"stack is empty\"")
	}
}

func TestTopEmptyStack(t *testing.T) {
	s := MakeStack[int](0)
	s.Push(9)
	s.Pop()
	v, err := s.Top()

	if v != nil {
		t.Fatalf("Top() return not nil")
	}
	if err == nil || err.Error() != "stack is empty" {
		t.Fatalf("Top() = nil, nil, want match for nil, \"stack is empty\"")
	}
}
