// The Maybe type encapsulates an optional value.
// A value of type Maybe a either contains a value of type a (represented as Just a),
// or it is empty (represented as Nothing).
// Using Maybe is a good way to deal with errors or exceptional cases without resorting to drastic measures such as error.
package maybe

// Maybe encapsulates an optional value.
type Maybe[T any] interface {
	IsJust() bool
	Get() T
}

// Just represents non empty case of Maybe
type Just[T any] struct {
	obj T
}

func (j Just[T]) IsJust() bool {
	return true
}

func (n Just[T]) Get() T {
	return n.obj
}

func NewJust[T any](v T) Just[T] {
	return Just[T]{
		obj: v,
	}
}

// Nothing represents empty case of Maybe
type Nothing[T any] struct{}

func NewNothing[T any]() Nothing[T] {
	return Nothing[T]{}
}

func (n Nothing[T]) IsJust() bool {
	return false
}

func (n Nothing[T]) Get() T {
	panic("maybe.Nothing.Get(): cannot get nothing")
}

// Maybe_ takes a default value, a function, and a Maybe value.
// If the Maybe value is Nothing, the function returns the default value.
// Otherwise, it applies the function to the value inside the Just and returns the result.
//func Maybe_[A any, B any](_default B, f func(a A) B, m Maybe[A]) B {
//
//}
