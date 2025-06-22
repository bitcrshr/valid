package validators

import "fmt"

type pointerValidator[T any, V Validator[T]] struct {
	checks        []func(*T) error
	elemValidator V
}

var _ PointerValidator[int, NumberValidator[int]] = &pointerValidator[int, NumberValidator[int]]{}

func NewPointerValidator[T any, V Validator[T]](elemValidator V) PointerValidator[T, V] {
	return &pointerValidator[T, V]{
		checks:        make([]func(*T) error, 0),
		elemValidator: elemValidator,
	}
}

func (v *pointerValidator[T, V]) Validate(value *T) error {
	for _, check := range v.checks {
		if err := check(value); err != nil {
			return err
		}
	}

	if value != nil {
		if err := v.elemValidator.Validate(*value); err != nil {
			return err
		}
	}

	return nil
}

func (v *pointerValidator[T, V]) ValidateAny(value any) error {
	t, ok := value.(*T)
	if !ok {
		return fmt.Errorf("expected value of type %T, but got %T", t, value)
	}

	return v.Validate(t)
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
