package function

// Curry binds an input parameter to a function that has two arguments
// and returns a new function that only has one. Achieves the same things
// as a lambda that closes over a variable.
// For more details see https://en.wikipedia.org/wiki/Currying
// Same as Curry1of2
func Curry[A any, B any, C any](f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

func Curry1[A any, B any](f func(A) B, a A) func() B {
	return func() B {
		return f(a)
	}
}

// Curry1of2, read: Curry first of two arguments
// Same as Curry
func Curry1of2[A any, B any, C any](f func(A, B) C, a A) func(B) C {
	return Curry(f, a)
}

// Curry2of2, read: Curry the second of two arguments
func Curry2of2[A any, B any, C any](f func(A, B) C, b B) func(A) C {
	return func(a A) C {
		return f(a, b)
	}
}

// Curry2of2E, read: Curry the second of two arguments for a function that also
// returns an error
func Curry2of2E[A any, B any, C any](f func(A, B) (C, error), b B) func(A) (C, error) {
	return func(a A) (C, error) {
		return f(a, b)
	}
}

func Curry2[A any, B any, C any](f func(A, B) C, a A, b B) func() C {
	return func() C {
		return f(a, b)
	}
}

func Curry2E[A any, B any, C any](f func(A, B) (C, error), a A, b B) func() (C, error) {
	return func() (C, error) {
		return f(a, b)
	}
}

func Panic(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

func Panic1[A any](f func() (A, error)) A {
	a, err := f()
	if err != nil {
		panic(err)
	}
	return a
}
