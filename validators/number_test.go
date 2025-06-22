package validators_test

import (
	"math"
	"testing"

	"github.com/bitcrshr/valid/validators"
	"golang.org/x/exp/constraints"
)

type numTestCase[T constraints.Integer | constraints.Float] struct {
	num  T
	pass bool
}

type numTest[T constraints.Integer | constraints.Float] struct {
	v     validators.NumberValidator[T]
	cases []numTestCase[T]
}

func TestNumberValidatorInt(t *testing.T) {
	intTests := []numTest[int]{
		{
			v: validators.NewNumberValidator[int]().Positive(),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: false},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().Negative(),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().Zero(),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: true},
				{num: -2, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().NonZero(),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: false},
				{num: -2, pass: true},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().LT(10),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: 10, pass: false},
				{num: 2000, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().LTE(10),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: 10, pass: true},
				{num: 2000, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().GT(10),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: false},
				{num: -2, pass: false},
				{num: 10, pass: false},
				{num: 2000, pass: true},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().GTE(10),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: false},
				{num: -2, pass: false},
				{num: 10, pass: true},
				{num: 2000, pass: true},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().EqualTo(10),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: false},
				{num: -2, pass: false},
				{num: 10, pass: true},
				{num: 2000, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().NotEqualTo(10),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: 10, pass: false},
				{num: 2000, pass: true},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().In(1, 2, 3),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: false},
				{num: -2, pass: false},
				{num: 1, pass: true},
				{num: 2, pass: true},
				{num: 3, pass: true},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().In(),
			cases: []numTestCase[int]{
				{num: 5, pass: false},
				{num: 0, pass: false},
				{num: -2, pass: false},
				{num: 1, pass: false},
				{num: 2, pass: false},
				{num: 3, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: false},
			},
		},

		{
			v: validators.NewNumberValidator[int]().NotIn(1, 2, 3),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: 1, pass: false},
				{num: 2, pass: false},
				{num: 3, pass: false},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().NotIn(),
			cases: []numTestCase[int]{
				{num: 5, pass: true},
				{num: 0, pass: true},
				{num: -2, pass: true},
				{num: 1, pass: true},
				{num: 2, pass: true},
				{num: 3, pass: true},
				{num: math.MaxInt, pass: true},
				{num: math.MinInt, pass: true},
			},
		},

		{
			v: validators.NewNumberValidator[int]().
				Positive().
				LT(42).
				GTE(13).
				In(20, 21, 500),
			cases: []numTestCase[int]{
				{num: 13, pass: false},
				{num: 20, pass: true},
				{num: 21, pass: true},
				{num: 500, pass: false},
				{num: -1, pass: false},
				{num: math.MaxInt, pass: false},
				{num: math.MinInt, pass: false},
			},
		},
	}

	for _, test := range intTests {
		for _, c := range test.cases {
			err := test.v.Validate(c.num)

			if err != nil && c.pass {
				t.Errorf("expected %v to pass", c.num)
			}

			if err == nil && !c.pass {
				t.Errorf("expected %v to fail", c.num)
			}
		}
	}
}
