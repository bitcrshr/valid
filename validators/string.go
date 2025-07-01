package validators

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type stringValidator[T ~string] struct {
	*baseValidator[T, StringValidator[T]]
}

var _ StringValidator[string] = &stringValidator[string]{}

func NewStringValidator[T ~string]() StringValidator[T] {
	v := &stringValidator[T]{}
	v.baseValidator = newBaseValidator[T, StringValidator[T]](v)

	return v
}

func (v *stringValidator[T]) Empty() StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if len(t) != 0 {
				return fmt.Errorf("expected `%s` to not be empty", string(t))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotEmpty() StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if len(t) == 0 {
				return fmt.Errorf("expected `%s` to not be empty", string(t))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) Len(l int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if len(t) != l {
				return fmt.Errorf("expected `%s` to have len %d, but got %d", string(t), l, len(t))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) MinLen(min int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if len(t) < min {
				return fmt.Errorf("expected `%s` to have min len %d, but got %d", string(t), min, len(t))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) MaxLen(max int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if len(t) > max {
				return fmt.Errorf("expected `%s` to have max len %d, but got %d", string(t), max, len(t))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) EqualTo(other T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t != other {
				return fmt.Errorf("expected `%s` to equal `%s`", string(t), string(other))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotEqualTo(other T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if t == other {
				return fmt.Errorf("expected `%s` not to equal `%s`", string(t), string(other))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) HasPrefix(prefix T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !strings.HasPrefix(string(t), string(prefix)) {
				return fmt.Errorf("expected `%s` to have prefix `%s`", string(t), string(prefix))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotHasPrefix(prefix T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.HasPrefix(string(t), string(prefix)) {
				return fmt.Errorf("expected `%s` not to have prefix `%s`", string(t), string(prefix))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) HasSuffix(suffix T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !strings.HasSuffix(string(t), string(suffix)) {
				return fmt.Errorf("expected `%s` to have suffix `%s`", string(t), string(suffix))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotHasSuffix(suffix T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.HasSuffix(string(t), string(suffix)) {
				return fmt.Errorf("expected `%s` not to have suffix `%s`", string(t), string(suffix))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) Contains(needle T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !strings.Contains(string(t), string(needle)) {
				return fmt.Errorf("expected `%s` to contain `%s`", string(t), string(needle))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotContains(needle T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.Contains(string(t), string(needle)) {
				return fmt.Errorf("expected `%s` not to contain `%s`", string(t), string(needle))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) ContainsAtLeast(needle T, count int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.Count(string(t), string(needle)) < count {
				return fmt.Errorf("expected `%s` to contain at least %d instances of `%s`", string(t), count, string(needle))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) ContainsAtMost(needle T, count int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.Count(string(t), string(needle)) > count {
				return fmt.Errorf("expected `%s` to contain at most %d instances of `%s`", string(t), count, string(needle))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) ContainsExact(needle T, count int) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if strings.Count(string(t), string(needle)) != count {
				return fmt.Errorf("expected `%s` to contain exactly %d instances of `%s`", string(t), count, string(needle))
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) In(haystack ...T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !slices.Contains(haystack, t) {
				return fmt.Errorf("expected `%s` to be in (%#v)", string(t), haystack)
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotIn(haystack ...T) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if slices.Contains(haystack, t) {
				return fmt.Errorf("expected `%s` not to be in (%#v)", string(t), haystack)
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) Matches(regex *regexp.Regexp) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if !regex.MatchString(string(t)) {
				return fmt.Errorf("expected `%s` to match regex `%s`", string(t), regex.String())
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) NotMatches(regex *regexp.Regexp) StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if regex.MatchString(string(t)) {
				return fmt.Errorf("expected `%s` not to match regex `%s`", string(t), regex.String())
			}

			return nil
		},
	)

	return v
}

func (v *stringValidator[T]) ValidUUID() StringValidator[T] {
	v.checks = append(
		v.checks,
		func(t T) error {
			if _, err := uuid.Parse(string(t)); err != nil {
				return fmt.Errorf("expected `%s` to be a valid uuid: %v", string(t), err)
			}

			return nil
		},
	)

	return v
}
