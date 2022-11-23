package typing

// Optional is a semantic marker that a value is optional and might not be set.
type Optional[T any] struct {
	// Value holds a pointer to the value if the value is set, or nil if the value is not set.
	//
	// This a temporary solution:
	// Currently Go does not support both: `type Optional[T any] T` (see https://github.com/golang/go/issues/45639)
	// and embedded fields out of type parameters.
	// Because of that on serializations it leads to an extra layer/level.
	Value *T
}

// Set sets the value.
func (o *Optional[T]) Set(v T) {
	o.Value = &v
}

// Unset unsets the value
func (o *Optional[T]) Unset() {
	o.Value = nil
}

// IsSet returns true if a value is set.
func (o Optional[T]) IsSet() bool {
	return o.Value != nil
}

// GetOrZero returns the value if it is set or returns the zero value.
func (o Optional[T]) GetOrZero() T {
	if o.IsSet() {
		return o.Get()
	}
	var zeroValue T
	return zeroValue
}

// Get returns the value (or panics if it is not set).
func (o Optional[T]) Get() T {
	return *o.Value
}

// Opt returns an optional value which holds the provided value
func Opt[T any](value T) Optional[T] {
	return Optional[T]{Value: &value}
}
