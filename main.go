package main

import (
	"fmt"
	"iter"
	"slices"
)

/* Iterator:

- function that produces values one at a time, without creating the entire collection up front.
- Magic happens through "yield"

 -> Generates value on demand (lazy evaluation)

	Lazy evaluation is a strategy that delays calculating an expression's value until it's actually needed,
	avoiding unnecessary computations and saving resources.
*/

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](src []V) *Iterator[V] {
	return &Iterator[V]{
		seq: slices.Values(src),
	}
}

func (it *Iterator[V]) Filter(pred func(V) bool) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !pred(v) {
					continue
				}
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Collect() []V {
	return slices.Collect(it.seq)
}

func (it *Iterator[V]) ForEach(f func(V)) {
	for v := range it.seq {
		f(v)
	}
}

func main() {

	result := From([]int{1, 2, 3, 4, 5, 6, 7, 8}).
		Filter(func(x int) bool {
			return x%2 == 0
		}).
		Map(func(x int) int {
			return x * 10
		}).
		Collect()

	fmt.Println(result)

	From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool {
			return x >= 3
		}).
		Map(func(x int) int {
			return x * 2
		}).
		ForEach(func(x int) {
			fmt.Printf("%d ", x)
		})
}
