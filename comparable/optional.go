package comparable

import (
	"github.com/xaionaro-go/typing"
)

type Optional[T comparable] typing.Optional[T]

// Set sets the value.
func (o *Optional[T]) Set(v T) {
	(*typing.Optional[T])(o).Set(v)
}

// Unset unsets the value
func (o *Optional[T]) Unset() {
	(*typing.Optional[T])(o).Unset()
}

// IsSet returns true if a value is set.
func (o Optional[T]) IsSet() bool {
	return typing.Optional[T](o).IsSet()
}

// GetOrZero returns the value if it is set or returns the zero value.
func (o Optional[T]) GetOrZero() T {
	return typing.Optional[T](o).GetOrZero()
}

// Get returns the value (or panics if it is not set).
func (o Optional[T]) Get() T {
	return typing.Optional[T](o).Get()
}

// Equal returns true if `o` and `cmp` holds optional values.
//
// The `IsSet` boolean is also expected to match.
func (o Optional[T]) Equal(cmp Optional[T]) bool {
	if o.IsSet() != cmp.IsSet() {
		return false
	}
	if !o.IsSet() {
		return true
	}

	return o.Get() == cmp.Get()
}

// Opt returns an optional value which holds the provided value
func Opt[T comparable](value T) Optional[T] {
	return Optional[T]{Value: &value}
}
