package coll

import "cmp"

func MaxIndex[T cmp.Ordered](a []T) int {
	if len(a) == 0 {
		return -1
	}
	maxIdx := 0
	for i, v := range a {
		if v > a[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Sum[T int16 | int8 | int | int64](a []T) int64 {
	var s int64 = 0
	for _, v := range a {
		s += int64(v)
	}
	return s
}

func Map[T any, V any](a []T, translate func(T) V) []V {
	mapped := make([]V, len(a))
	for i, aa := range a {
		mapped[i] = translate(aa)
	}
	return mapped
}
