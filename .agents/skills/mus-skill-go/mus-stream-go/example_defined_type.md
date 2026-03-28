## Defined Type Example

```go
type Foo int


var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, w mus.Writer) (n int, err error) {
	return varint.Int.Marshal(int(v), w)
}

func (s fooMUS) Unmarshal(r mus.Reader) (v Foo, n int, err error) {
	tmp, n, err := varint.Int.Unmarshal(r)
	if err != nil {
		return
	}
	v = Foo(tmp)
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	return varint.Int.Size(int(v))
}

func (s fooMUS) Skip(r mus.Reader) (n int, err error) {
	return varint.Int.Skip(r)
}
```
