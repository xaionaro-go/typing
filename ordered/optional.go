package ordered

import (
	"github.com/xaionaro-go/typing/comparable"
	"golang.org/x/exp/constraints"
)

type Optional[T constraints.Ordered] comparable.Optional[T]

// Set sets the value.
func (o *Optional[T]) Set(v T) {
	(*comparable.Optional[T])(o).Set(v)
}

// Unset unsets the value
func (o *Optional[T]) Unset() {
	(*comparable.Optional[T])(o).Unset()
}

// IsSet returns true if a value is set.
func (o Optional[T]) IsSet() bool {
	return comparable.Optional[T](o).IsSet()
}

// GetOrZero returns the value if it is set or returns the zero value.
func (o Optional[T]) GetOrZero() T {
	return comparable.Optional[T](o).GetOrZero()
}

// Get returns the value (or panics if it is not set).
func (o Optional[T]) Get() T {
	return comparable.Optional[T](o).Get()
}

// Equal returns true if `o` and `cmp` holds optional values.
//
// The `IsSet` boolean is also expected to match.
func (o Optional[T]) Equal(cmp Optional[T]) bool {
	return comparable.Optional[T](o).Equal(comparable.Optional[T](cmp))
}

// Compare returns `-1` if `o` is less than `cmp`, `1` -- if greater, and `0` -- is equal
func (o Optional[T]) Compare(cmp Optional[T]) int {
	aIsSet, bIsSet := o.IsSet(), cmp.IsSet()
	if aIsSet != bIsSet {
		if aIsSet {
			return 1
		}
		return -1
	}
	if !aIsSet {
		return 0
	}

	a, b := o.Get(), cmp.Get()
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

// Opt returns an optional value which holds the provided value
func Opt[T constraints.Ordered](value T) Optional[T] {
	return Optional[T]{Value: &value}
}
