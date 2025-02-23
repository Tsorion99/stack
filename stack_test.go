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
		t.Fatal("Top() return not nil")
	}
	if err == nil || err.Error() != "stack is empty" {
		t.Fatal("Top() = nil, nil, want match for nil, \"stack is empty\"")
	}
}

func TestTopEmptyStack(t *testing.T) {
	s := MakeStack[int](0)
	s.Push(9)
	s.Pop()
	v, err := s.Top()

	if v != nil {
		t.Fatal("Top() return not nil")
	}
	if err == nil || err.Error() != "stack is empty" {
		t.Fatal("Top() = nil, nil, want match for nil, \"stack is empty\"")
	}
}

func TestLen(t *testing.T) {
	s := MakeStack[int](0)
	if len := s.Len(); len != 0 {
		t.Fatalf("Len() = %d, want match for 0", len)
	}

	s.Push(9, 6, 3)
	if len := s.Len(); len != 3 {
		t.Fatalf("Len() = %d, want match for 3", len)
	}

	s.Pop()
	if len := s.Len(); len != 2 {
		t.Fatalf("Len() = %d, want match for 2", len)
	}
}

func TestEmpty(t *testing.T) {
	s := MakeStack[int](0)
	if e := s.Empty(); e != true {
		t.Fatalf("Empty() = %t, want match for true", e)
	}

	s.Push(9, 6)
	if e := s.Empty(); e != false {
		t.Fatalf("Empty() = %t, want match for false", e)
	}

	s.Pop()
	if e := s.Empty(); e != true {
		t.Fatalf("Empty() = %t, want match for true", e)
	}
}
