# User Hints

With hints user can specify the serialization options for the type.

## Outer Hints

[Outer Hints](./outer.md) these hints could be applied to the user-provided type.

## Inner Hints

[Inner Hints](./inner.md) these hints could be applied to the struct field type 
or underlying type.

## Validator Functions

If the validator hint refers to a function (e.g., `func ValidateXYZ(T) error`,
see [Validation](./validation.md)), the generator MUST automatically wrap it in
`com.ValidatorFn[T]` to be used with the `mus` options. In other case it must
use the hint as is.

Example:

```go
// mus:lenVl = ValidateLength
type Foo string

func ValidateLength(l int) error {
	return nil
}

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewValidStringSer(
	stropts.WithLenValidator(com.ValidatorFn[int](ValidateLength)),
)
```