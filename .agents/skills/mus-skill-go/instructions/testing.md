# Testing

For each generated serializer, create a unit test in `mus.ai.gen_test.go` file 
using `github.com/mus-format/mus-go/test` package.

Example:

```go
func TestFoo(t *testing.T) {
	var (
		ser   = FooMUS
		cases = []Foo{
			// generate test cases
		}
	)
	test.Test(cases, ser, t)
	test.TestSkip(cases, ser, t)
}
```

> [!IMPORTANT]
> **Nil vs Empty Collections**: In test cases, for `array`, `slice`, and `map` 
> types, DO NOT use `nil`. Instead, always initialize them with an empty value 
> (e.g., `[]T{}` or `map[K]V{}`). This is because the MUS unmarshaler returns 
> empty collections instead of `nil`.

For each validator, create a separate test function:

```go
func TestFoo_`validator name`(t *testing.T) {	
	var (
		ser      = FooMUS
		testCase = Foo{...}
		wantErr  = errors.New("...")
	)
	test.TestValidation(testCase, ser, wantErr, t)
}
```
