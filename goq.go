// Package goq implements an unbounded, in-memory queue
// Queue never blocks on Enqueue, and will block on Dequeue
// until a value is enqueued, at which time it will immediately
// return that value.
//
// Quick start:
//
//    q := NewQueue[string]()
//    go func() {
//        for {
//            fmt.Println(q.Dequeue())
//        }
//    }()
//    q.Enqueue("internet")
//
package goq

import (
	"sync"
)

// Queue is a type that implements an unbounded queue
type Queue[T any] struct {
	q         *[]T
	writeCond *sync.Cond
}

// NewQueue returns a new, empty queue that can contain type T
func NewQueue[T any]() *Queue[T] {
	queue := make([]T, 0, 32)
	q := &Queue[T]{
		q:         &queue,
		writeCond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}

// ToSlice copies the elements in the queue to a new slice in a thread-safe manner
func (q *Queue[T]) ToSlice() []T {
	q.writeCond.L.Lock()
	defer q.writeCond.L.Unlock()
	res := make([]T, len(*q.q), len(*q.q))
	copy(res, *q.q)
	return res
}

func (q *Queue[T]) Enqueue(x T) {
	q.writeCond.L.Lock()
	defer q.writeCond.L.Unlock()
	*q.q = append(*q.q, x)
	// Dequeue might be waiting, notify it
	q.writeCond.Broadcast()
}

func (q *Queue[T]) Dequeue() T {
	q.writeCond.L.Lock()
	defer q.writeCond.L.Unlock()
	for q.empty() {
		q.writeCond.Wait()
	}
	x := (*q.q)[0]
	*q.q = (*q.q)[1:]
	return x
}

func (q *Queue[T]) empty() bool {
	return len(*q.q) == 0
}
