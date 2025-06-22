package validators_test

import (
	"testing"

	"github.com/bitcrshr/valid/validators"
)

func TestStructValidator(t *testing.T) {
	type Foo struct {
		Bar string
		Baz int
	}

	v := validators.NewStructValidator[Foo](validators.StructShape{
		"Bar": validators.NewStringValidator[string]().Contains("ooga"),
		"Baz": validators.NewNumberValidator[int]().Positive(),
	})

	foo1 := Foo{Bar: "ooga booga", Baz: 42}
	foo2 := Foo{Bar: "", Baz: 42}
	foo3 := Foo{Bar: "ooga booga", Baz: -1}
	foo4 := Foo{}

	if err := v.Validate(foo1); err != nil {
		t.Errorf("expected %#v to pass", foo1)
	}

	if err := v.Validate(foo2); err == nil {
		t.Errorf("expected %#v to fail", foo2)
	}

	if err := v.Validate(foo3); err == nil {
		t.Errorf("expected %#v to fail", foo3)
	}

	if err := v.Validate(foo4); err == nil {
		t.Errorf("expected %#v to fail", foo4)
	}
}
