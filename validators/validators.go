package validators

import (
	"cmp"
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
		// Satisfies(predicate func(T) error) NumberValidator[T]
	}

	MapValidator[K comparable, V any] interface {
		Validator[map[K]V]

		HasKey(key K) MapValidator[K, V]
		NotHasKey(key K) MapValidator[K, V]
		HasKeyIn(haystack ...K) MapValidator[K, V]
		NotHasKeyIn(haystack ...K) MapValidator[K, V]
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
	}

	CustomComparableSliceValidator[S ~[]E, E any, V Validator[E], F func(a, b E) bool] interface {
		SliceValidator[S, E, V]

		Contains(needle E) CustomComparableSliceValidator[S, E, V, F]
		NotContains(needle E) CustomComparableSliceValidator[S, E, V, F]
		ContainsAtLeast(needle E, count int) CustomComparableSliceValidator[S, E, V, F]
		ContainsAtMost(needle E, count int) CustomComparableSliceValidator[S, E, V, F]
		ContainsExact(needle E, count int) CustomComparableSliceValidator[S, E, V, F]
	}

	ComparableSliceValidator[S ~[]E, E comparable, V Validator[E]] interface {
		SliceValidator[S, E, V]

		Contains(needle E) ComparableSliceValidator[S, E, V]
		NotContains(needle E) ComparableSliceValidator[S, E, V]
		ContainsAtLeast(needle E, count int) ComparableSliceValidator[S, E, V]
		ContainsAtMost(needle E, count int) ComparableSliceValidator[S, E, V]
		ContainsExact(needle E, count int) ComparableSliceValidator[S, E, V]
	}

	CustomOrderedSliceValidator[S ~[]E, E any, V Validator[E], F func(a, b E) int] interface {
		CustomComparableSliceValidator[S, E, V, func(a E, b E) bool]

		Sorted() CustomOrderedSliceValidator[S, E, V, F]
		Unsorted() CustomOrderedSliceValidator[S, E, V, F]
	}

	OrderedSliceValidator[S ~[]E, E cmp.Ordered, V Validator[E]] interface {
		ComparableSliceValidator[S, E, V]

		Sorted() OrderedSliceValidator[S, E, V]
		Unsorted() OrderedSliceValidator[S, E, V]
	}

	PointerValidator[T any, V Validator[T]] interface {
		Validator[*T]

		ElemValidator() V

		Nil() PointerValidator[T, V]
		NotNil() PointerValidator[T, V]
	}
)
