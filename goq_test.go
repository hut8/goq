package goq

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hello")
	q.Enqueue("world")
	// select {
	// case q.Enqueue <- "hello":
	// 	t.Log("enqueue")
	// default:
	// 	t.Errorf("enqueue would block")
	// }
	// select {
	// case q.Enqueue <- "world":
	// 	t.Log("enqueue")
	// default:
	// 	t.Errorf("enqueue would block")
	// }
	if len(*q.q) != 2 {
		t.Errorf("expected 2 elements in queue, got %v",
			len(*q.q))
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hello")
	q.Enqueue("world")
	if len(*q.q) != 2 {
		t.Errorf("expected 2 elements in queue, got %v",
			len(*q.q))
	}
	hello := q.Dequeue()
	if len(*q.q) != 1 {
		t.Errorf("expected 1 element in queue, got %v",
			len(*q.q))
	}
	world := q.Dequeue()
	if len(*q.q) != 0 {
		t.Errorf("expected empty queue, got %v",
			len(*q.q))
	}
	if hello != "hello" {
		t.Errorf("expected hello, got %v", hello)
	}
	if world != "world" {
		t.Errorf("expected world, got %v", world)
	}
}

func TestToSlice(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hello")
	q.Enqueue("world")
	s := q.ToSlice()
	if len(s) != 2 {
		t.Errorf("expected slice of size 2, got %v",
			len(s))
	}
}
