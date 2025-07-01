package validators

import "fmt"

type pointerValidator[T any, V Validator[T]] struct {
	*baseValidator[*T, PointerValidator[T, V]]
	elemValidator V
}

var _ PointerValidator[int, NumberValidator[int]] = &pointerValidator[int, NumberValidator[int]]{}

func NewPointerValidator[T any, V Validator[T]](elemValidator V) PointerValidator[T, V] {
	v := &pointerValidator[T, V]{
		elemValidator: elemValidator,
	}
	v.baseValidator = newBaseValidator[*T, PointerValidator[T, V]](v)

	return v
}

func (v *pointerValidator[T, V]) Nil() PointerValidator[T, V] {
	v.checks = append(
		v.checks,
		func(t *T) error {
			if t != nil {
				return fmt.Errorf("expected %#v to be nil", t)
			}

			return nil
		},
	)

	return v
}

func (v *pointerValidator[T, V]) NotNil() PointerValidator[T, V] {
	v.checks = append(
		v.checks,
		func(t *T) error {
			if t == nil {
				return fmt.Errorf("expected %#v to not be nil", t)
			}

			return nil
		},
	)

	return v
}

func (v *pointerValidator[T, V]) ElemValidator() V {
	return v.elemValidator
}
