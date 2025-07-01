package validators

import "fmt"

type baseValidator[T any, Super Validator[T]] struct {
	checks []func(T) error
	super  Super
}

func newBaseValidator[T any, Super Validator[T]](super Super) *baseValidator[T, Super] {
	return &baseValidator[T, Super]{
		checks: make([]func(T) error, 0),
		super:  super,
	}
}

func (v *baseValidator[T, Super]) Validate(value T) error {
	for _, check := range v.checks {
		if err := check(value); err != nil {
			return err
		}
	}

	return nil
}

func (v *baseValidator[T, Super]) ValidateAny(value any) error {
	t, ok := value.(T)
	if !ok {
		return fmt.Errorf("expected value of type %T, but found %T", t, value)
	}

	return v.Validate(t)
}

func (v *baseValidator[T, Super]) Satisfies(check func(T) error) Super {
	v.checks = append(v.checks, check)
	return v.super
}
