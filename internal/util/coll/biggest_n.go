package coll

import (
	"slices"

	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/constraints"
)

type BiggestN[T any, P constraints.Ordered] struct {
	cap int
	pq  *lane.PriorityQueue[T, P] // max-oriented
}

func NewBiggestN[T any, P constraints.Ordered](n int) *BiggestN[T, P] {
	if n <= 0 {
		panic("NewBiggestN: capacity must be > 0")
	}
	return &BiggestN[T, P]{
		cap: n,
		pq:  lane.NewMinPriorityQueue[T, P](),
	}
}

// Len returns the current number of stored elements (<= capacity).
func (s *BiggestN[T, P]) Len() int { return int(s.pq.Size()) }

// Capacity returns the configured maximum number of elements.
func (s *BiggestN[T, P]) Capacity() int { return s.cap }

// Push inserts a value with its priority. If the size exceeds capacity,
// the current largest-priority element will be evicted. The method returns
// whether an eviction occurred.
// UNSAFE: cannot be used in concurrent contexts.
func (s *BiggestN[T, P]) Push(v T, p P) (evicted bool) {
	s.pq.Push(v, p)
	if int(s.pq.Size()) > s.cap {
		// Pop removes the max (since this is a max-oriented PQ)
		_, _, _ = s.pq.Pop()
		return true
	}
	return false
}

// Pop removes and returns the current largest element among the kept ones
// (i.e., the worst among the N smallest). It returns ok=false if empty.
func (s *BiggestN[T, P]) Pop() (v T, p P, ok bool) {
	return s.pq.Pop()
}

// PeekWorst returns, without removing, the current largest priority element
// among the kept ones. It returns ok=false if empty.
func (s *BiggestN[T, P]) PeekWorst() (v T, p P, ok bool) {
	return s.pq.Head()
}

// Empty reports whether no elements are stored.
func (s *BiggestN[T, P]) Empty() bool { return s.pq.Empty() }

func (s *BiggestN[T, P]) PopAllDecending() []T {
	tmp := s.PopAllAccending()
	slices.Reverse(tmp)
	return tmp
}

func (s *BiggestN[T, P]) PopAllAccending() []T {
	values := make([]T, 0, s.Len())
	for {
		v, _, ok := s.pq.Pop()
		if !ok {
			break
		}
		values = append(values, v)
	}
	return values
}
