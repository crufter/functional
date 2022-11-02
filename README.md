# functional

A functional "standard" library for Golang.

## Why

- To spark discussion and ideas
- To make algorithmic/data oriented work with Go easier and faster
- To push the edges of Go's type system (test driving generics)
- For fun (putting the fun in functional)
- Why not?

8 years ago I wrote an [article](https://crufter.com/everyday-hassles-in-go) about the limitations of the typesystem of Go. It resonated with a lot of people so much that it even ended up on [Wikipedia](<https://en.wikipedia.org/wiki/Go_(programming_language)>) as a source (under the criticism section - seems like I'm a professional hater!).

Since the majority of the points in my post were about the lack of generics, it was about time I test drove the generic features.

This project is that test drive.

## How

Currently this package is extremely heavily inspired by the Haskell standard library. To the point of stealing their comments and examples. Don't tell them. I hope their stuff is MIT :fingers-crossed:.

## Examples

From the obvious ones like Map

```go
double := func(i int) int {
	return i * 2
}
v := list.Map(double, []int{1, 2, 3})
fmt.Println(v)
// prints 2, 4, 6
```

to more peculiar ones like

```go
list.StripPrefix([]string{"f", "o", "o"}, []string{"f", "o", "o", "b", "a", "r"})
// returns maybe.NewJust([]string{"b", "a", "r"})

list.StripPrefix([]string{"f", "o", "o"}, []string{"b", "a", "r", "f", "o", "o"})
// returns maybe.NewNothing[[]string]()
```

or fairly crazy stuff like [Currying](https://en.wikipedia.org/wiki/Currying)

```go
add := func(i, j int) int {
	return i + j
}
addTo3 := function.Curry(add, 3)
addTo3(5)
// pints 8
```

to [tuples](https://en.wikipedia.org/wiki/Tuple)

```go
// Uncons decomposes a list into its head and tail.
// Uncons[A any](xs []A) maybe.Maybe[tuple.Couple[A, []A]]
m := list.Uncons([]int{1, 2, 3})
if !m.IsJust() {
    panic("uncons returned Nothing")
}
t := m.Get()
head := t.E1()
tail := t.E2()
fmt.Printn(head, tail)
// prints 1, []int{2, 3}
```

have fun and don't blame me.

## Status

Extremely early.
