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

func Filter[T constraints.Ordered](elements []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[T, C any](input []T, f func(T) C) []C {
	res := make([]C, 0)
	for _, v := range input {
		res = append(res, f(v))
	}
	return res
}
