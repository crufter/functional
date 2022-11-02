package function_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/crufter/functional/data/function"
	"github.com/stretchr/testify/assert"
)

func TestCurry(t *testing.T) {
	print := func(i string, j int) string {
		return fmt.Sprintf("%v%v", i, j)
	}
	p := function.Curry(print, "hi")
	assert.Equal(t, "hi5", p(5))
}

func TestCurry0(t *testing.T) {
	double := func(i int) int {
		return i * 2
	}
	double5 := function.Curry1(double, 5)
	assert.Equal(t, 10, double5())
}

func TestCurry2(t *testing.T) {
	add := func(i, j int) int {
		return i + j
	}
	add3And4 := function.Curry2(add, 3, 4)
	assert.Equal(t, 7, add3And4())
}

func TestCurry1of2(t *testing.T) {
	print := func(i string, j int) string {
		return fmt.Sprintf("%v%v", i, j)
	}
	p := function.Curry1of2(print, "hi")
	assert.Equal(t, "hi5", p(5))
}

func TestCurry2of2(t *testing.T) {
	print := func(i string, j int) string {
		return fmt.Sprintf("%v%v", i, j)
	}
	p := function.Curry2of2(print, 5)
	assert.Equal(t, "hi5", p("hi"))
}

func TestPanic(t *testing.T) {
	print := func(i string, j int) (string, error) {
		if j == 6 {
			return "", errors.New("baad")
		}
		return fmt.Sprintf("%v%v", i, j), nil
	}

	printHi5 := function.Curry2E(print, "hi", 5)
	assert.Equal(t, "hi5", function.Panic1(printHi5))
}
