package validators_test

import (
	"regexp"
	"testing"

	"github.com/bitcrshr/valid/validators"
	"github.com/google/uuid"
)

func TestStringValidator(t *testing.T) {
	cases := []struct {
		v     validators.StringValidator[string]
		tests []struct {
			str  string
			pass bool
		}
	}{
		{
			v: validators.NewStringValidator[string]().Empty(),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "", pass: true},
				{str: "foo", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotEmpty(),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "", pass: false},
				{str: "foo", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().Len(5),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "hello", pass: true},
				{str: "foo", pass: false},
				{str: "", pass: false},
				{str: "ohgreatheavens", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().MinLen(2),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "hello", pass: true},
				{str: "foo", pass: true},
				{str: "", pass: false},
				{str: "ohgreatheavens", pass: true},
				{str: "E", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().MaxLen(2),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "hello", pass: false},
				{str: "foo", pass: false},
				{str: "", pass: true},
				{str: "ohgreatheavens", pass: false},
				{str: "EE", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().EqualTo("yeet"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "yeet", pass: true},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotEqualTo("yeet"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "yeet", pass: false},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().HasPrefix("oog"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: true},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotHasPrefix("oog"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: false},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().HasSuffix("ooga"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: true},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotHasSuffix("ooga"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: false},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().Contains("a b"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: true},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotContains("a b"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga", pass: false},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().ContainsAtLeast("a b", 2),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga b", pass: true},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().ContainsAtMost("a b", 2),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga b", pass: true},
				{str: "ooga booga", pass: true},
				{str: "ooga booga b a b", pass: false},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().ContainsExact("a b", 2),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "ooga booga b", pass: true},
				{str: "ooga booga", pass: false},
				{str: "ooga booga b a b", pass: false},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().In("abc", "def", "ghi"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "abc", pass: true},
				{str: "", pass: false},
				{str: "yoink", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotIn("abc", "def", "ghi"),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "jkl", pass: true},
				{str: "abc", pass: false},
				{str: "", pass: true},
				{str: "yoink", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().Matches(regexp.MustCompile("^[a-z]+$")),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "jkl", pass: true},
				{str: "abc", pass: true},
				{str: "", pass: false},
				{str: "yoink", pass: true},
				{str: "8675309", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().NotMatches(regexp.MustCompile("^[a-z]+$")),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: "jkl", pass: false},
				{str: "abc", pass: false},
				{str: "", pass: true},
				{str: "yoink", pass: false},
				{str: "8675309", pass: true},
			},
		},

		{
			v: validators.NewStringValidator[string]().ValidUUID(),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: uuid.NewString(), pass: true},
				{str: uuid.Nil.String(), pass: true},
				{str: "abc", pass: false},
				{str: "", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().
				MinLen(3).
				MaxLen(10).
				NotContains("foo").
				In("bar", "baz", "bazinga").
				NotEmpty().
				HasPrefix("ba").
				NotHasSuffix("az"),

			tests: []struct {
				str  string
				pass bool
			}{
				{str: "bar", pass: true},
				{str: "baz", pass: false},
				{str: "bazinga", pass: true},
				{str: "", pass: false},
				{str: "hello, world!", pass: false},
			},
		},

		{
			v: validators.NewStringValidator[string]().ValidUUID(),
			tests: []struct {
				str  string
				pass bool
			}{
				{str: uuid.NewString(), pass: true},
				{str: uuid.Nil.String(), pass: true},
				{str: "abc", pass: false},
				{str: "", pass: false},
			},
		},
	}

	for _, c := range cases {
		for _, test := range c.tests {
			err := c.v.Validate(test.str)

			if test.pass && err != nil {
				t.Errorf("Expected `%s` to pass", test.str)
			}

			if !test.pass && err == nil {
				t.Errorf("Expected `%s` to fail", test.str)
			}
		}
	}
}
