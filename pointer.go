package pointer

import "reflect"

// Ref returns a pointer to the given value.
func Ref[T any](t T) *T {
	return &t
}

// Deref dereference the given pointer.
func Deref[T any](t *T) T {
	return *t
}

// SafeDeref dereference the given pointer and return the zero value if the pointer is nil.
func SafeDeref[T any](t *T) T {
	if t == nil {
		return *new(T)
	}
	return *t
}

// Real will deference the `any` value until the actual base type
// is reached.
func Real(value any) any {
	if value == nil {
		return nil
	}
	rt := reflect.TypeOf(value)
	switch {
	// Non pointer; return given value.
	case rt.Kind() != reflect.Ptr:
		return value
	// Pointer; deref until we're at the bottom.
	default:
		rv := reflect.ValueOf(value)
		for rv.Kind() == reflect.Ptr && !rv.IsNil() {
			rv = rv.Elem()
		}
		return rv.Interface()
	}
}
