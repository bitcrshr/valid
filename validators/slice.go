package validators

import "fmt"

type sliceValidator[S ~[]E, E any, V Validator[E]] struct {
	*baseValidator[S, SliceValidator[S, E, V]]
	elemValidator V
}

var _ SliceValidator[[]string, string, StringValidator[string]] = &sliceValidator[[]string, string, StringValidator[string]]{}

func NewSliceValidator[S ~[]E, E any, V Validator[E]](elemValidator V) SliceValidator[S, E, V] {
	v := &sliceValidator[S, E, V]{
		elemValidator: elemValidator,
	}
	v.baseValidator = newBaseValidator[S, SliceValidator[S, E, V]](v)

	v.checks = append(
		v.checks,
		func(s S) error {
			for i, el := range s {
				if err := elemValidator.Validate(el); err != nil {
					return fmt.Errorf("element at index %d did not satisfy elem validator: %v", i, err)
				}
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) Empty() SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			if len(s) != 0 {
				return fmt.Errorf("expected %v to be empty", s)
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) NotEmpty() SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			if len(s) == 0 {
				return fmt.Errorf("expected %v not to be empty", s)
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) Len(l int) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			if len(s) == l {
				return fmt.Errorf("expected %v to have len %d", s, l)
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) MinLen(min int) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			if len(s) < min {
				return fmt.Errorf("expected %v to have min len %d", s, min)
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) MaxLen(max int) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			if len(s) > max {
				return fmt.Errorf("expected %v to have max len %d", s, max)
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) AllSatisfy(validator V) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			for i, el := range s {
				if err := validator.Validate(el); err != nil {
					return fmt.Errorf("element at index %d did not satisfy validator: %v", i, err)
				}
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) AnySatisfy(validator V) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			for _, el := range s {
				if err := validator.Validate(el); err == nil {
					return nil
				}
			}

			return fmt.Errorf("expected at least one element in %v to pass validator", s)
		},
	)
	return v
}

func (v *sliceValidator[S, E, V]) NoneSatisfy(validator V) SliceValidator[S, E, V] {
	v.checks = append(
		v.checks,
		func(s S) error {
			for i, el := range s {
				if err := validator.Validate(el); err == nil {
					return fmt.Errorf("element at index %d passed validator: %v", i, err)
				}
			}

			return nil
		},
	)

	return v
}

func (v *sliceValidator[S, E, V]) ElemValidator() V {
	return v.elemValidator
}
