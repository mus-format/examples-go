# Outer Hints

These hints could be applied to the user-provided type.

## Serializer Name
 
With the `mus:name` hint, the user can specify a custom name for the 
generated MUS serializer.
 
```go
// mus:name = CustomSerializer
type Foo struct {
  ...
}
```
 
Example:
 
```go
// mus:name = CustomFoo
type Foo struct {
 	num int
}
 
var CustomFooMUS = customFooMUS{}
 
type customFooMUS struct{}
 
func (s customFooMUS) Marshal( ... ) (n int) { ... }
...
```

## Validator

```go
// mus:vl = ...
T
```

Where:

- vl: `com.Validator[T]` interface implementation

Example:

```go
// mus:vl = ValidateFoo
type Foo int

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v1, n, err = varint.Int.Unmarshal(...)
	if err != nil {
		return
	}
	v = Foo(v1)
	if err = ValidateFoo(v); err != nil {
		return
	}
	return
}
...
```

## Interface Implementations

With the `mus:impls` hint, the user can specify the types that implement an 
interface.

```go
// mus:impls = Foo, Bar, Baz
```

Example:

```go
// mus:impls = Foo, Bar
type MyInterface interface {
  Method()
}
```

The generator will use this hint to create a serializer for the interface type.

## Marshaller Interface

With the `mus:marshaller` hint, the user can specify that the generated 
interface serializer should use `mus.MarshallerTyped` interface.

## Generated Code Location

With the `mus:path` hint, the user can specify the location of the serializer. 
This hint is used for both:
1. Specifying the location (package) where the code should be generated.
2. Specifying the location where the generator should look for a previously 
   generated serializer.

```go
// mus:path = package_name
T
```

Example:

```go
// mus:path = github.com/user/project/pkg
type Foo struct {
	...
}
```