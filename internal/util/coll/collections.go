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

func Fill[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}

func NewSlice[T any](size int, val T) []T {
	slice := make([]T, size)
	Fill(slice, val)
	return slice
}
