package valid

import (
	"github.com/bitcrshr/valid/validators"
	"golang.org/x/exp/constraints"
)

func String() validators.StringValidator[string] {
	return validators.NewStringValidator[string]()
}

func StringLike[T ~string]() validators.StringValidator[T] {
	return validators.NewStringValidator[T]()
}

func Int() validators.NumberValidator[int] {
	return validators.NewNumberValidator[int]()
}

func Int8() validators.NumberValidator[int8] {
	return validators.NewNumberValidator[int8]()
}

func Int16() validators.NumberValidator[int16] {
	return validators.NewNumberValidator[int16]()
}

func Int32() validators.NumberValidator[int32] {
	return validators.NewNumberValidator[int32]()
}

func Int64() validators.NumberValidator[int64] {
	return validators.NewNumberValidator[int64]()
}

func Uint() validators.NumberValidator[uint] {
	return validators.NewNumberValidator[uint]()
}

func Uint8() validators.NumberValidator[uint8] {
	return validators.NewNumberValidator[uint8]()
}

func Uint16() validators.NumberValidator[uint16] {
	return validators.NewNumberValidator[uint16]()
}

func Uint32() validators.NumberValidator[uint32] {
	return validators.NewNumberValidator[uint32]()
}

func Uint64() validators.NumberValidator[uint64] {
	return validators.NewNumberValidator[uint64]()
}

func Uintptr() validators.NumberValidator[uintptr] {
	return validators.NewNumberValidator[uintptr]()
}

func Float32() validators.NumberValidator[float32] {
	return validators.NewNumberValidator[float32]()
}

func Float64() validators.NumberValidator[float64] {
	return validators.NewNumberValidator[float64]()
}

func Numeric[T constraints.Integer | constraints.Float]() validators.NumberValidator[T] {
	return validators.NewNumberValidator[T]()
}

func Pointer[T any, V validators.Validator[T]](elemValidator V) validators.PointerValidator[T, V] {
	return validators.NewPointerValidator(elemValidator)
}

func Slice[S ~[]E, E any, V validators.Validator[E]](elemValidator V) validators.SliceValidator[S, E, V] {
	return validators.NewSliceValidator[S](elemValidator)
}

func Struct[T any](shape validators.StructShape) *validators.StructValidator[T] {
	return validators.NewStructValidator[T](shape)
}
