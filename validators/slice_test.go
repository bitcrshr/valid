package validators_test

import (
	"testing"

	"github.com/bitcrshr/valid/validators"
)

func TestSliceValidator(t *testing.T) {
	v := validators.NewSliceValidator[[]string](
		validators.NewStringValidator[string]().
			NotEmpty().
			MinLen(2),
	).
		MaxLen(3).
		NotEmpty().
		NoneSatisfy(
			validators.NewStringValidator[string]().
				EqualTo("nope"),
		)

	s := []string{"hello"}
	if err := v.Validate(s); err != nil {
		t.Errorf("expected %v to pass", s)
	}

	s = []string{"", "a", "yoyo"}
	if err := v.Validate(s); err == nil {
		t.Errorf("expected %v to fail", s)
	}
}
