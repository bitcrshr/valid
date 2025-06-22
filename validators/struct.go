package validators

import (
	"fmt"
	"reflect"
)

type StructShape map[string]AnyValidator

type StructValidator[T any] struct {
	checks []func(T) error
	shape  StructShape
}

func NewStructValidator[T any](shape StructShape) *StructValidator[T] {
	return &StructValidator[T]{
		shape: shape,
	}
}

func (v *StructValidator[T]) Validate(value T) error {
	for fieldName, validator := range v.shape {
		field := reflect.
			ValueOf(value).
			FieldByName(fieldName)

		if err := validator.ValidateAny(field.Interface()); err != nil {
			return err
		}
	}

	return nil
}

func (v *StructValidator[T]) ValidateAny(value any) error {
	t, ok := value.(T)
	if !ok {
		return fmt.Errorf("expected value to be of type %T, but got %T", value, t)
	}

	return v.Validate(t)
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
