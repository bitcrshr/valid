package validators_test

import (
	"testing"

	"github.com/bitcrshr/valid/validators"
)

func ptrTo[T any](v T) *T {
	return &v
}

func TestPointerValidator(t *testing.T) {
	v := validators.NewPointerValidator(
		validators.NewNumberValidator[int]().In(1, 2, 3).Positive(),
	).NotNil()

	if err := v.Validate(ptrTo(2)); err != nil {
		t.Errorf("expected %#v to pass", ptrTo(2))
	}

	if err := v.Validate(ptrTo(5)); err == nil {
		t.Errorf("expected %#v to fail", ptrTo(5))
	}

	if err := v.Validate(nil); err == nil {
		t.Errorf("expected %#v to fail", nil)
	}

	v2 := validators.NewPointerValidator(
		validators.NewStringValidator[string]().Contains("foo"),
	)

	if err := v2.Validate(ptrTo("foobar")); err != nil {
		t.Errorf("expected %s to pass", "foobar")
	}

	if err := v2.Validate(ptrTo("")); err == nil {
		t.Errorf("expected %q to fail", "")
	}

	if err := v2.Validate(nil); err != nil {
		t.Errorf("expected %#v to pass", nil)
	}

	v3 := validators.NewPointerValidator(
		validators.NewNumberValidator[float64]().EqualTo(3.14),
	).Nil()

	if err := v3.Validate(nil); err != nil {
		t.Errorf("expected %#v to pass", nil)
	}

	if err := v3.Validate(ptrTo(8.675309)); err == nil {
		t.Errorf("expected %#v to fail", ptrTo(7.675309))
	}
}
