package validators

import (
	"golang.org/x/exp/constraints"
	"regexp"
)

type (
	AnyValidator interface {
		ValidateAny(value any) error
	}
	Validator[T any] interface {
		AnyValidator

		Validate(value T) error
	}

	StringValidator[T ~string] interface {
		Validator[T]

		Empty() StringValidator[T]
		NotEmpty() StringValidator[T]
		Len(l int) StringValidator[T]
		MinLen(min int) StringValidator[T]
		MaxLen(max int) StringValidator[T]
		EqualTo(other T) StringValidator[T]
		NotEqualTo(other T) StringValidator[T]
		HasPrefix(prefix T) StringValidator[T]
		NotHasPrefix(prefix T) StringValidator[T]
		HasSuffix(suffix T) StringValidator[T]
		NotHasSuffix(suffix T) StringValidator[T]
		Contains(needle T) StringValidator[T]
		NotContains(needle T) StringValidator[T]
		ContainsAtLeast(needle T, count int) StringValidator[T]
		ContainsAtMost(needle T, count int) StringValidator[T]
		ContainsExact(needle T, count int) StringValidator[T]
		In(haystack ...T) StringValidator[T]
		NotIn(haystack ...T) StringValidator[T]
		Matches(regex *regexp.Regexp) StringValidator[T]
		NotMatches(regex *regexp.Regexp) StringValidator[T]
		ValidUUID() StringValidator[T]

		Satisfies(check func(T) error) StringValidator[T]
	}

	NumberValidator[T constraints.Integer | constraints.Float] interface {
		Validator[T]

		Positive() NumberValidator[T]
		Negative() NumberValidator[T]
		Zero() NumberValidator[T]
		NonZero() NumberValidator[T]
		LT(upper T) NumberValidator[T]
		LTE(upper T) NumberValidator[T]
		GT(lower T) NumberValidator[T]
		GTE(lower T) NumberValidator[T]
		EqualTo(other T) NumberValidator[T]
		NotEqualTo(other T) NumberValidator[T]
		In(haystack ...T) NumberValidator[T]
		NotIn(haystack ...T) NumberValidator[T]

		Satisfies(check func(T) error) NumberValidator[T]
	}

	MapValidator[K comparable, V any] interface {
		Validator[map[K]V]

		Empty() MapValidator[K, V]
		NotEmpty() MapValidator[K, V]
		HasKey(key K) MapValidator[K, V]
		NotHasKey(key K) MapValidator[K, V]
		HasKeyIn(haystack ...K) MapValidator[K, V]
		NotHasKeyIn(haystack ...K) MapValidator[K, V]

		Satisfies(check func(map[K]V) error) MapValidator[K, V]
	}

	SliceValidator[S ~[]E, E any, V Validator[E]] interface {
		Validator[S]

		ElemValidator() V

		Empty() SliceValidator[S, E, V]
		NotEmpty() SliceValidator[S, E, V]
		Len(l int) SliceValidator[S, E, V]
		MinLen(min int) SliceValidator[S, E, V]
		MaxLen(max int) SliceValidator[S, E, V]
		AllSatisfy(v V) SliceValidator[S, E, V]
		AnySatisfy(v V) SliceValidator[S, E, V]
		NoneSatisfy(v V) SliceValidator[S, E, V]

		Satisfies(check func(S) error) SliceValidator[S, E, V]
	}

	PointerValidator[T any, V Validator[T]] interface {
		Validator[*T]

		ElemValidator() V

		Nil() PointerValidator[T, V]
		NotNil() PointerValidator[T, V]

		Satisfies(check func(*T) error) PointerValidator[T, V]
	}
)
