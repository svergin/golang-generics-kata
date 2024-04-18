package generic

import "golang.org/x/exp/constraints"

func ForEach[T any](elements []T, f func(T) T) {
	for i, e := range elements {
		elements[i] = f(e)
	}
}

func Collect[T any](element T) (res []T) {
	// TODO Implement me!
	c := make(chan T)
	//res := make([]T,0)
	c <- element
	go func() {
		for v := range c {
			res = append(res, v)
		}
		close(c)
	}()
	return res
}

// returns a new datatype that contains only those elements matching a given predicate function
// Works!
func Filter[T constraints.Ordered](elements []T, matches func(T) bool, otherOps ...func(T) any) []T {
	res := make([]T, 0)
	for _, v := range elements {
		if matches(v) {
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
