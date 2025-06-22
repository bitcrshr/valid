# valid

*valid* is a validation library for Go, with a focus on:
* extendability
* simplicity
* speed
* versatility

One of the main differentiating factors for `valid` as compared to alternatives is it is just as useful for types you *don't own* as ones you do. This makes it very easy to work with generated code that you don't wish to modify yourself.


>  ⚠️ **BEWARE!** This is very much a WIP, and is not yet ready for production use. 

### Overview

All validators in this library build upon the `Validator` interface:

```go
type AnyValidator interface {
    Validate(value any) error
}

type Validator[T any] interface {
    AnyValidator
    
	Validate(value T) error
}
```

You can use the built-in validators:

```go
package valid_demo

import (
	"regexp"

	"github.com/bitcrshr/valid"
	"github.com/bitcrshr/valid/validators"
)

type User struct {
	Id   string
	Name string
	Age  int
}

// It is generally recommended that you define your validator once and reuse
// it throughout your codebase.
//
// Types whose validation is well-defined an unlikely to change can be hidden
// behind the validators.Validator interface to denote that this validator is
// "final" and shouldn't be extended.
func UserValidator() validators.Validator[*User] {
	return valid.Pointer(
		valid.Struct[User](
			validators.StructShape{
				"Id": valid.String().NotEmpty().ValidUUID(),
				"Name": valid.String().
					MinLen(5).
					Matches(regexp.MustCompile("^[a-zA-Z-]*$")),
				"Age": valid.Int().GT(21).LT(150),
			},
		),
	).NotNil()
}

// Alternatively, you can return the root validator so this can serve as your
// "base" and be extended where needed.
func BaseUserValidator() validators.PointerValidator[User, *validators.StructValidator[User]] {
	return valid.Pointer(
		valid.Struct[User](
			validators.StructShape{
				"Id": valid.String().NotEmpty().ValidUUID(),
			},
		),
	).NotNil()
}

func main() {
	if err := UserValidator().Validate(&User{
		Id:   "0e49b3e4-77ea-4c89-bdba-64a7d4efd042",
		Name: "Bobby Axelrod",
		Age:  43,
	}); err != nil {
		panic(err)
	}

	userExt := BaseUserValidator()
	userExt.ElemValidator().Shape()["Name"] = valid.String().NotEqualTo("Banned User")

	if err := userExt.Validate(nil); err != nil {
		panic(err)
	}
}
```

Or, you can create your own and implement validation logic however you like:

```go
package valid_demo

import (
	"fmt" 

	"github.com/bitcrshr/valid/validators"
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	Name string
	Age int
}

type userValidator struct {}

func (uv *userValidator) Validate(user *User) error {
	switch {
	case user == nil:
		return fmt.Errorf("user cannot be nil")
	case user.Id == uuid.Nil:
		return fmt.Errorf("user.Id cannot be nil uuid")
	case user.Name == "":
		return fmt.Errorf("user.Name cannot be empty")
	case user.Age < 21 || user.Age > 150: 
		return fmt.Errorf("user.Age must be in [21, 150]")
	default:
		return nil
	}
}

func (uv *userValidator) ValidateAny(maybeUser any) error {
	user, ok := maybeUser.(*User)
	if !ok {
		return fmt.Errorf("expected value of type *User, but found %T", maybeUser)
	}

	return uv.Validate(user)
}

func NewUserValidator() validators.Validator[*User] {
	return &userValidator{}
}
```

### Goals / Roadmap

- [ ] Provide optional mechanisms to avoid or minimize performance hit of reflection for structs
	- [ ] Code generation in the spirit of [ent](https://github.com/ent/ent)
	- [ ] Caching of reflection data
	- [ ] ???
- [ ] Benchmarks in comparison with popular alternatives
- [ ] Custom errors (to support things like gRPC errors, custom messages, etc.)
- [ ] Tooling for convenient usage in tests``
