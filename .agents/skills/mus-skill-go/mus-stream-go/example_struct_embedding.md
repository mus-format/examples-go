# Struct Embedding Example

```go
type Bar struct {
	num int
}

type Foo struct {
	Bar
	str string
}

var BarMUS = barMUS{}

type barMUS struct{}

func (s barMUS) Marshal(v Bar, w mus.Writer) (n int, err error) {
	return varint.Int.Marshal(v.num, w)
}

func (s barMUS) Unmarshal(r mus.Reader) (v Bar, n int, err error) {
	v.num, n, err = varint.Int.Unmarshal(r)
	return
}

func (s barMUS) Size(v Bar) (size int) {
	return varint.Int.Size(v.num)
}

func (s barMUS) Skip(r mus.Reader) (n int, err error) {
	n, err = varint.Int.Skip(r)
	return
}

var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, w mus.Writer) (n int, err error) {
	n, err = BarMUS.Marshal(v.Bar, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Marshal(v.str, w)
	n += n1
	return
}

func (s fooMUS) Unmarshal(r mus.Reader) (v Foo, n int, err error) {
	v.Bar, n, err = BarMUS.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal(r)
	n += n1
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	size = BarMUS.Size(v.Bar)
	return size + ord.String.Size(v.str)
}

func (s fooMUS) Skip(r mus.Reader) (n int, err error) {
	n, err = BarMUS.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(r)
	n += n1
	return
}
```
