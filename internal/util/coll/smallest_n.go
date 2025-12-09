package coll

import (
	"slices"

	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/constraints"
)

// SmallestN keeps at most N elements with the smallest priorities.
//
// It is implemented on top of lane.NewMaxPriorityQueue: the queue head is the
// current largest priority among the kept elements. When pushing past the
// capacity, the largest element is evicted so only the smallest N remain.
//
// Type parameters:
//   - T: stored value type
//   - P: priority type (must be ordered)
type SmallestN[T any, P constraints.Ordered] struct {
	cap int
	pq  *lane.PriorityQueue[T, P] // max-oriented
}

// NewSmallestN creates a new container that retains only the N smallest
// elements by priority.
func NewSmallestN[T any, P constraints.Ordered](n int) *SmallestN[T, P] {
	if n <= 0 {
		panic("NewSmallestN: capacity must be > 0")
	}
	return &SmallestN[T, P]{
		cap: n,
		pq:  lane.NewMaxPriorityQueue[T, P](),
	}
}

// Len returns the current number of stored elements (<= capacity).
func (s *SmallestN[T, P]) Len() int { return int(s.pq.Size()) }

// Capacity returns the configured maximum number of elements.
func (s *SmallestN[T, P]) Capacity() int { return s.cap }

// Push inserts a value with its priority. If the size exceeds capacity,
// the current largest-priority element will be evicted. The method returns
// whether an eviction occurred.
// UNSAFE: cannot be used in concurrent contexts.
func (s *SmallestN[T, P]) Push(v T, p P) (evicted bool) {
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
func (s *SmallestN[T, P]) Pop() (v T, p P, ok bool) {
	return s.pq.Pop()
}

// PeekWorst returns, without removing, the current largest priority element
// among the kept ones. It returns ok=false if empty.
func (s *SmallestN[T, P]) PeekWorst() (v T, p P, ok bool) {
	return s.pq.Head()
}

// Empty reports whether no elements are stored.
func (s *SmallestN[T, P]) Empty() bool { return s.pq.Empty() }

// PopAllAccending pops all values ordered by increasing
// priority (smallest priority first).
func (s *SmallestN[T, P]) PopAllAccending() []T {
	tmp := make([]T, 0, s.Len())
	for {
		v, _, ok := s.pq.Pop()
		if !ok {
			break
		}
		tmp = append(tmp, v)
	}
	slices.Reverse(tmp)
	return tmp
}
