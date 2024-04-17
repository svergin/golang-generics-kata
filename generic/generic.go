package generic

import "golang.org/x/exp/constraints"

func Filter2[T constraints.Ordered](elements []T, f func(T) bool, f2 ...func(T) any) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Filter[T constraints.Ordered](elements []T, filter func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[T, C any](input []T, mapping func(T) C) []C {
	res := make([]C, 0)
	for _, v := range input {
		res = append(res, mapping(v))
	}
	return res
}

func Reduce[T, R any](elements []T, aggregate func(T, R) R) R {
	var result R
	for _, v := range elements {
		result = aggregate(v, result)
	}

	return result
}
