# Defined Type Example

```go
type Foo int

var FooMUS = fooMUS{}

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
	return varint.Int.Marshal(int(v), bs)
}

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
	tmp, n, err := varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	v = Foo(tmp)
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	return varint.Int.Size(int(v))
}

func (s fooMUS) Skip(bs []byte) (n int, err error) {
	return varint.Int.Skip(bs)
}
```
