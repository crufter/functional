package list_test

import (
	"testing"

	"github.com/crufter/functional/data/list"
	"github.com/crufter/functional/data/maybe"
	"github.com/crufter/functional/data/tuple"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	assert.Equal(t, list.Append([]int{1, 2}, []int{3, 4}), []int{1, 2, 3, 4})

	a := []int{1, 2, 3}
	b := []int{4, 5}
	v := list.Append(a, b)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, v)
	assert.Equal(t, []int{1, 2, 3}, a)
	assert.Equal(t, []int{4, 5}, b)
}

func TestHead(t *testing.T) {
	a := []int{1, 2, 3}
	v := list.Head(a)
	assert.Equal(t, 1, v)
	assert.Equal(t, []int{1, 2, 3}, a)

	t.Run("test empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		a := []int{}
		list.Head(a)
	})
}

func TestLast(t *testing.T) {
	a := []int{1, 2, 3}
	v := list.Last(a)
	assert.Equal(t, 3, v)
	assert.Equal(t, []int{1, 2, 3}, a)

	t.Run("test empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		list.Last([]int{})
	})
}

func TestInit(t *testing.T) {
	assert.Equal(t, list.Init([]int{1, 2, 3}), []int{1, 2})

	a := []int{1, 2, 3}
	v := list.Init(a)
	assert.Equal(t, []int{1, 2}, v)
	assert.Equal(t, []int{1, 2, 3}, a)

	t.Run("test empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		list.Init([]int{})
	})
}

func TestTail(t *testing.T) {
	assert.Equal(t, list.Tail([]int{1, 2, 3}), []int{2, 3})
	assert.Equal(t, list.Tail([]int{1}), []int{})

	a := []int{1, 2, 3}
	v := list.Init(a)
	assert.Equal(t, []int{1, 2}, v)
	assert.Equal(t, []int{1, 2, 3}, a)

	t.Run("test empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		list.Init([]int{})
	})
}

func TestUncons(t *testing.T) {
	assert.Equal(t, list.Uncons([]int{}), maybe.NewNothing[tuple.Couple[int, []int]]())
	assert.Equal(t, list.Uncons([]int{1}), maybe.NewJust(tuple.NewCouple(1, []int{})))
	assert.Equal(t, list.Uncons([]int{1, 2, 3}), maybe.NewJust(tuple.NewCouple(1, []int{2, 3})))

	ret := list.Uncons([]int{1, 2, 3})
	assert.Equal(t, true, ret.IsJust())
	assert.Equal(t, 1, ret.Get().E1())
	assert.Equal(t, []int{2, 3}, ret.Get().E2())

	ret = list.Uncons([]int{})
	assert.Equal(t, false, ret.IsJust())

}

func TestSingleton(t *testing.T) {
	assert.Equal(t, list.Singleton(true), []bool{true})
}

func TestMap(t *testing.T) {
	double := func(i int) int {
		return i * 2
	}
	assert.Equal(t, list.Map(double, []int{1, 2, 3}), []int{2, 4, 6})
	assert.Equal(t, list.Map(double, []int{}), []int{})
}

func TestReverse(t *testing.T) {
	assert.Equal(t, list.Reverse([]int{}), []int{})
	assert.Equal(t, list.Reverse([]int{42}), []int{42})
	assert.Equal(t, list.Reverse([]int{2, 5, 7}), []int{7, 5, 2})
}

func TestIntersperse(t *testing.T) {
	assert.Equal(t, list.Intersperse("-", []string{"h", "e", "l", "l", "o"}), []string{"h", "-", "e", "-", "l", "-", "l", "-", "o"})
	assert.Equal(t, list.Intersperse(1, []int{}), []int{})
	assert.Equal(t, list.Intersperse(1, []int{2}), []int{2})
}

func TestTranspose(t *testing.T) {
	assert.Equal(t, list.Transpose([][]int{{1, 2, 3}, {4, 5, 6}}), [][]int{{1, 4}, {2, 5}, {3, 6}})
	assert.Equal(t, list.Transpose([][]int{{10, 11}, {20}, {}, {30, 31, 32}}), [][]int{{10, 20, 30}, {11, 31}, {32}})
}

//func TestSubsquences(t *testing.T) {
//	assert.Equal(t, list.Subsequences([]string{"a", "b", "c"}), [][]string{{""}, {"a"}, {"b"}, {"ab"}, {"c"}, {"ac"}, {"bc"}, {"abc"}})
//}

func TestPermutations(t *testing.T) {
	assert.Equal(t, list.Permutations([]string{"a", "b", "c"}), [][]string{{"a", "b", "c"}, {"b", "a", "c"}, {"c", "b", "a"}, {"b", "c", "a"}, {"c", "a", "b"}, {"a", "c", "b"}})
}

func TestConcat(t *testing.T) {
	assert.Equal(t, list.Concat([][]int{{1, 2, 3}, {4, 5}, {6}, {}}), []int{1, 2, 3, 4, 5, 6})
}

func TestStripPrefix(t *testing.T) {
	assert.Equal(t, list.StripPrefix([]string{"f", "o", "o"}, []string{"f", "o", "o", "b", "a", "r"}), maybe.NewJust([]string{"b", "a", "r"}))
	assert.Equal(t, list.StripPrefix([]string{"f", "o", "o"}, []string{"f", "o", "o"}), maybe.NewJust([]string{}))
	assert.Equal(t, list.StripPrefix([]string{"f", "o", "o"}, []string{"b", "a", "r", "f", "o", "o"}), maybe.NewNothing[[]string]())
	assert.Equal(t, list.StripPrefix([]string{"f", "o", "o"}, []string{"b", "a", "r", "f", "o", "o", "b", "a", "z"}), maybe.NewNothing[[]string]())
}
