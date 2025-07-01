package validators

import (
	"fmt"
	"reflect"
)

type StructShape map[string]AnyValidator

type StructValidator[T any] struct {
	*baseValidator[T, *StructValidator[T]]
	shape StructShape
}

func NewStructValidator[T any](shape StructShape) *StructValidator[T] {
	v := &StructValidator[T]{
		shape: shape,
	}
	v.baseValidator = newBaseValidator(v)

	return v
}

func (v *StructValidator[T]) Zero() *StructValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !reflect.ValueOf(t).IsZero() {
				return fmt.Errorf("expected %v to be zero value of %T", t, t)
			}

			return nil
		},
	)

	return v
}

func (v *StructValidator[T]) NotZero() *StructValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if reflect.ValueOf(t).IsZero() {
				return fmt.Errorf("expected %v to not be zero value of %T", t, t)
			}

			return nil
		},
	)

	return v
}

func (v *StructValidator[T]) Shape() StructShape {
	return v.shape
}
