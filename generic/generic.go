package generic

import "golang.org/x/exp/constraints"

func Filter[T constraints.Ordered](elements []T, f func(T) bool, otherOps ...func(T) any) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[T, C any](input []T, mapping func(T) C, otherOps ...func(T) any) []C {
	res := make([]C, 0)
	for _, v := range input {
		res = append(res, mapping(v))
	}
	return res
}

func Reduce[T, R any](elements []T, aggregate func(T, R) R, otherOps ...func(T) any) R {
	var result R
	for _, v := range elements {
		result = aggregate(v, result)
	}

	return result
}
