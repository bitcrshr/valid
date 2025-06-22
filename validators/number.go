package validators

import (
	"fmt"
	"slices"

	"golang.org/x/exp/constraints"
)

type numberValidator[T constraints.Integer | constraints.Float] struct {
	checks []func(T) error
}

func NewNumberValidator[T constraints.Integer | constraints.Float]() NumberValidator[T] {
	return &numberValidator[T]{
		checks: make([]func(T) error, 0),
	}
}

var _ NumberValidator[int] = NewNumberValidator[int]()

func (v *numberValidator[T]) Validate(value T) error {
	for _, check := range v.checks {
		if err := check(value); err != nil {
			return err
		}
	}

	return nil
}

func (v *numberValidator[T]) ValidateAny(value any) error {
	t, ok := value.(T)
	if !ok {
		return fmt.Errorf("expected value to be of type %T, but got %T", value, t)
	}

	return v.Validate(t)
}

func (v *numberValidator[T]) Positive() NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t < 0 {
				return fmt.Errorf("expected %v to be positive", t)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) Negative() NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t > 0 {
				return fmt.Errorf("expected %v to be negative", t)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) Zero() NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t != 0 {
				return fmt.Errorf("expected %v to be zero", t)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) NonZero() NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t == 0 {
				return fmt.Errorf("expected %v to be nonzero", t)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) LT(upper T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t >= upper {
				return fmt.Errorf("expected %v to be less than %v", t, upper)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) LTE(upper T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t > upper {
				return fmt.Errorf("expected %v to be less than or equal to %v", t, upper)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) GT(lower T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t <= lower {
				return fmt.Errorf("expected %v to be greater than %v", t, lower)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) GTE(lower T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t < lower {
				return fmt.Errorf("expected %v to be greater than or equal to %v", t, lower)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) EqualTo(other T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t != other {
				return fmt.Errorf("expected %v to be equal to %v", t, other)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) NotEqualTo(other T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t == other {
				return fmt.Errorf("expected %v not to be equal to %v", t, other)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) In(haystack ...T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !slices.Contains(haystack, t) {
				return fmt.Errorf("expected %v to be in (%v)", t, haystack)
			}

			return nil
		},
	)

	return v
}

func (v *numberValidator[T]) NotIn(haystack ...T) NumberValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if slices.Contains(haystack, t) {
				return fmt.Errorf("expected %v not to be in (%v)", t, haystack)
			}

			return nil
		},
	)

	return v
}
