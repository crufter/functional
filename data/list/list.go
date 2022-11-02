package list

import (
	"reflect"

	"github.com/crufter/functional/data/maybe"
	"github.com/crufter/functional/data/tuple"
)

//
// Basic functions
//

// Append two lists, ie
//
//	list.Append([]int{1, 2}, []int{3, 4}) == []int{1, 2, 3, 4}
func Append[A any](xs []A, ys []A) []A {
	ret := make([]A, len(xs)+len(ys))
	for i, v := range xs {
		ret[i] = v
	}
	for i, v := range ys {
		ret[len(xs)+i] = v
	}
	return ret
}

// Head extracts the first element of a list, which must be non-empty.
// Only here for completeness' sake, use [0] instead.
//
//	list.Head([]int{1, 2, 3}) = []int{1}
//	list.Head([]int{}) = panic
func Head[A any](xs []A) A {
	if len(xs) == 0 {
		panic("list.Head: empty list argument")
	}
	return xs[0]
}

// Last extracts the last element of a list, which must be non-empty.
//
//	list.Last([]int{1, 2, 3}) = 3
//	list.Last([]int{}) = panic
func Last[A any](xs []A) A {
	if len(xs) == 0 {
		panic("list.Last: empty list argument")
	}

	return xs[len(xs)-1]
}

// Init O(n). Return all the elements of a list except the last one. The list must be non-empty.
//
//	list.Init([]int{1, 2, 3}) = []int{1, 2}
//	list.Init([]int{}) = panic
func Init[A any](xs []A) []A {
	if len(xs) == 0 {
		panic("list.Init: empty list argument")
	}
	ret := make([]A, len(xs)-1)
	for i, v := range xs {
		if i == len(xs)-1 {
			break
		}
		ret[i] = v
	}

	return ret
}

// Tail extracts the elements after the head of a list, which must be non-empty.
//
//	list.Tail([]int{1, 2, 3}) == []int{2,3}
//	list.Tail([]int{1}) == []int{}
//	list.Tail([]int{}) == panic
func Tail[A any](xs []A) []A {
	if len(xs) == 0 {
		panic("list.Tail: empty list argument")
	}
	ret := make([]A, len(xs)-1)
	for i, v := range xs[1:] {
		ret[i] = v
	}

	return ret
}

// Uncons decomposes a list into its head and tail.
//
//	list.Uncons([]int{}) == maybe.NewNothing[tuple.Couple[int, []int]]
//	list.Uncons([]int{1}) == maybe.NewJust(tuple.NewCouple(1, []int{}))
//	list.Uncons([]int{1, 2, 3}) == maybe.NewJust(tuple.NewCouple(1, []int{2, 3}))
func Uncons[A any](xs []A) maybe.Maybe[tuple.Couple[A, []A]] {
	if len(xs) == 0 {
		return maybe.NewNothing[tuple.Couple[A, []A]]()
	}
	return maybe.NewJust(tuple.NewCouple(xs[0], xs[1:]))
}

// Singleton produces a singleton list.
//
//	list.Singleton(true) == []bool{true}
func Singleton[T any](a T) []T {
	return []T{a}
}

//
// List transformations
//

// Map(f, xs) returns the list obtained by applying f to each element of xs, i.e.,
//
//	list.Map(func(i int) int {
//		return i * 2
//	}, []int{1, 2, 3})) == []int{2, 4, 6}
func Map[A any, B any](f func(A) B, xs []A) []B {
	ret := make([]B, len(xs))
	for i, v := range xs {
		ret[i] = f(v)
	}
	return ret
}

// Reverse returns the elements of xs in reverse order. xs must be finite.
//
//	list.Reverse([]int{}) == []int{}
//	list.Reverse([]int{42}) == []int{42}
//	list.Reverse([]int{2, 5, 7}) == []int{7, 5, 2}
func Reverse[A any](xs []A) []A {
	ret := make([]A, len(xs))
	for i, v := range xs {
		ret[len(xs)-1-i] = v
	}
	return ret
}

// Intersperse takes an element and a list and `intersperses' that element between the elements of the list. For example,
//
//	list.Intersperse("-", []string{"h", "e", "l", "l", "o"}) == []string{"h", "-", "e", "-", "l", "-", "l", "-", "o"}
func Intersperse[A any](x A, xs []A) []A {
	if len(xs) == 0 {
		return make([]A, 0)
	}
	ret := make([]A, len(xs)*2-1)
	for i, v := range xs {
		ret[i*2] = v
		if i*2 == len(ret)-1 {
			break
		}
		ret[i*2+1] = x
	}
	return ret
}

// @todo intercalate

// Transpose transposes the rows and columns of its argument. For example,
//
//	list.Transpose([][]int{{1, 2, 3}, {4, 5, 6}}) == [][]int{{1, 4}, {2, 5}, {3, 6}}
//	list.Transpose([][]int{{10, 11}, {20}, {}, {30, 31, 32}}) == [][]int{{10, 20, 30}, {11, 31}, {32}}
func Transpose[A any](xss [][]A) [][]A {
	// horribly inelegant and inefficient implementation
	max := 0
	for _, xs := range xss {
		if len(xs) > max {
			max = len(xs)
		}
	}
	ret := make([][]A, 0)
	for i := 0; i < max; i++ {
		tmp := make([]A, 0)
		for _, xs := range xss {
			if len(xs) > i {
				tmp = append(tmp, xs[i])
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

// Subsequences returns the list of all subsequences of the argument.
//
//	list.Subsequences([]string{"a", "b", "c"}) == [][]string{{""}, {"a"}, {"b"}, {"ab"}, {"c"}, {"ac"}, {"bc"}, {"abc"}}
//
// @todo not done
//func Subsequences[A any](xs []A) [][]A {
//	return nil
//}

// Permutations function returns the list of all permutations of the argument.
//
//	list.Permutations([]string{"a", "b", "c"}) == [][]string{{"a", "b", "c"},{"b", "a", "c"}, {"c", "b", "a"}, {"b", "c", "a"}, {"c", "a", "b"}, {"a", "c", "b"}}
func Permutations[A any](xs []A) [][]A {
	var helper func([]A, int)
	res := [][]A{}

	helper = func(arr []A, n int) {
		if n == 1 {
			tmp := make([]A, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(xs, len(xs))
	return res
}

// Concat returns the concatenation of all the elements of a container of lists.
//
//	list.Concat([][]int{{1, 2, 3}, {4, 5}, {6}, {}}) == []int{1, 2, 3, 4, 5, 6}
func Concat[A any](xss [][]A) []A {
	ret := make([]A, 0)
	for _, xs := range xss {
		ret = append(ret, xs...)
	}
	return ret
}

// StripPrefix function drops the given prefix from a list.
// It returns Nothing if the list did not start with the prefix given, or Just the list after the prefix, if it does.
//
//	list.StripPrefix([]string{"f", "o", "o"}, []string{"f", "o", "o", "b", "a", "r"}) == maybe.NewJust([]string{"b", "a", "r"})
//	list.StripPrefix([]string{"f", "o", "o"}, []string{"f", "o", "o"}) == maybe.NewJust([]string{})
//	list.StripPrefix([]string{"f", "o", "o"}, []string{"b", "a", "r", "f", "o", "o"}) == maybe.NewNothing[[]string]()
//	list.StripPrefix([]string{"f", "o", "o"}, []string{"b", "a", "r", "f", "o", "o", "b", "a", "z"}) == maybe.NewNothing[[]string]()
func StripPrefix[A any](prefix, ys []A) maybe.Maybe[[]A] {
	ret := make([]A, 0)
	for i, v := range ys {
		switch {
		// DeepEqual == unreal levels of cruft. enjoy
		case len(prefix) > i && reflect.DeepEqual(prefix[i], v):
			continue
		case len(prefix) > i:
			return maybe.NewNothing[[]A]()
		}
		ret = append(ret, v)
	}
	return maybe.NewJust(ret)
}
