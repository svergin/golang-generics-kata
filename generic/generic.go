package generic

import "golang.org/x/exp/constraints"

func Filter[T constraints.Ordered](elements []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
