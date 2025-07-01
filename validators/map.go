package validators

import "fmt"

type mapValidator[K comparable, V any] struct {
	*baseValidator[map[K]V, MapValidator[K, V]]
}

func NewMapValidator[K comparable, V any]() MapValidator[K, V] {
	v := &mapValidator[K, V]{}

	v.baseValidator = newBaseValidator[map[K]V, MapValidator[K, V]](v)

	return v
}

func (v *mapValidator[K, V]) Empty() MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			if len(m) > 0 {
				return fmt.Errorf("expected %v to be empty", m)
			}

			return nil
		},
	)

	return v
}

func (v *mapValidator[K, V]) NotEmpty() MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			if len(m) == 0 {
				return fmt.Errorf("expected %v not to be empty", m)
			}

			return nil
		},
	)

	return v
}

func (v *mapValidator[K, V]) HasKey(key K) MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			if _, ok := m[key]; !ok {
				return fmt.Errorf("expected %v to have key %v", m, key)
			}

			return nil
		},
	)

	return v
}

func (v *mapValidator[K, V]) NotHasKey(key K) MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			if _, ok := m[key]; ok {
				return fmt.Errorf("expected %v not to have key %v", m, key)
			}

			return nil
		},
	)

	return v
}

func (v *mapValidator[K, V]) HasKeyIn(haystack ...K) MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			for _, needle := range haystack {
				if _, ok := m[needle]; ok {
					return nil
				}
			}

			return fmt.Errorf("expected %v to have at least 1 key in (%v)", m, haystack)
		},
	)

	return v
}

func (v *mapValidator[K, V]) NotHasKeyIn(haystack ...K) MapValidator[K, V] {
	v.checks = append(
		v.checks,
		func(m map[K]V) error {
			for _, needle := range haystack {
				if _, ok := m[needle]; ok {
					return fmt.Errorf("expected %v not to have any keys in (%v)", m, haystack)
				}
			}

			return nil
		},
	)

	return v
}
