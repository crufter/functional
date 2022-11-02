// package tuple implements container types
// that are similar to fixed sized arrays but with
// elements that can differ in type.
// They could be also though of as structs without fields.
// see https://en.wikipedia.org/wiki/Tuple
package tuple

// Couple is like a fixed sized array but its elements can be of different types.
type Couple[A any, B any] struct {
	_1 A
	_2 B
}

func NewCouple[A any, B any](a A, b B) Couple[A, B] {
	return Couple[A, B]{
		_1: a,
		_2: b,
	}
}

// E1 returns the first element of a Couple
func (n Couple[A, B]) E1() A {
	return n._1
}

// E2 returns the second element of a Couple
func (n Couple[A, B]) E2() B {
	return n._2
}

// Triple is like a fixed sized array but its elements can be of different types.
type Triple[A any, B any, C any] struct {
	_1 A
	_2 B
	_3 C
}

func NewTriple[A any, B any, C any](a A, b B, c C) Triple[A, B, C] {
	return Triple[A, B, C]{
		_1: a,
		_2: b,
		_3: c,
	}
}

// E1 returns the first element of a Triple
func (n Triple[A, B, C]) E1() A {
	return n._1
}

// E2 returns the second element of a Triple
func (n Triple[A, B, C]) E2() B {
	return n._2
}

// E3 returns the second element of a Triple
func (n Triple[A, B, C]) E3() C {
	return n._3
}

// Quadruple is like a fixed sized array but its elements can be of different types.
type Quadruple[A any, B any, C any, D any] struct {
	_1 A
	_2 B
	_3 C
	_4 D
}

func NewQuadruple[A any, B any, C any, D any](a A, b B, c C, d D) Quadruple[A, B, C, D] {
	return Quadruple[A, B, C, D]{
		_1: a,
		_2: b,
		_3: c,
		_4: d,
	}
}

// E1 returns the first element of a Quadruple
func (n Quadruple[A, B, C, D]) E1() A {
	return n._1
}

// E2 returns the second element of a Quadruple
func (n Quadruple[A, B, C, D]) E2() B {
	return n._2
}

// E3 returns the second element of a Quadruple
func (n Quadruple[A, B, C, D]) E3() C {
	return n._3
}

// E4 returns the second element of a Quadruple
func (n Quadruple[A, B, C, D]) E4() D {
	return n._4
}
