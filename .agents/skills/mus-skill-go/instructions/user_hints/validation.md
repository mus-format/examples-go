## Validation

`Validator` interface is defined in the `github.com/mus-format/common-go` module.

```go
// Validator is the interface that wraps Validate method.
//
// Validate performs data validation.
type Validator[T any] interface {
	Validate(t T) error
}

// ValidatorFn is a functional implementation of the Validator interface.
type ValidatorFn[T any] func(t T) (err error)

func (fn ValidatorFn[T]) Validate(t T) (err error) {
	return fn(t)
}
```
