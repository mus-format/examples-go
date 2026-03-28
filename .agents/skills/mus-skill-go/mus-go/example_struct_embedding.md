# Struct Embedding Example

```go
type Bar struct {
	num int
}

type Foo struct {
	Bar
	str string
}

var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
	n = BarMUS.Marshal(v.Bar, bs)
	return n + ord.String.Marshal(v.str, bs[n:])
}

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
	// First field: direct assignment to n (optimization)
	v.Bar, n, err = BarMUS.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal(bs[n:])
	n += n1
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	size = BarMUS.Size(v.Bar)
	return size + ord.String.Size(v.str)
}

func (s fooMUS) Skip(bs []byte) (n int, err error) {
	// First field: direct assignment to n (optimization)
	n, err = BarMUS.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(bs[n:])
	n += n1
	return
}
```
