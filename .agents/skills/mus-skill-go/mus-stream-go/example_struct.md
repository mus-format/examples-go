# Struct Example

```go
type Foo struct {
	// mus:vl = ValidateNum
	num int 
	str string
	// mus:lenEnc = raw
	sl  []int 
}

var slicev1qt0ZcfFi3cECΔrX3nfjAΞΞ = ord.NewSliceSer[int](varint.Int, slops.WithLenSer[int](raw.Int))

var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, w mus.Writer) (n int, err error) {
	n, err = varint.Int.Marshal(v.num, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Marshal(v.str, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = slicev1qt0ZcfFi3cECΔrX3nfjAΞΞ.Marshal(v.sl, w)
	n += n1
	return
}

func (s fooMUS) Unmarshal(r mus.Reader) (v Foo, n int, err error) {
  // First field: direct assignment to n (optimization)
	v.num, n, err = varint.Int.Unmarshal(r)
	if err != nil {
		return
	}
	err = ValidateNum(v.num)
	if err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v.sl, n1, err = slicev1qt0ZcfFi3cECΔrX3nfjAΞΞ.Unmarshal(r)
	n += n1
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	size = varint.Int.Size(v.num)
	size += ord.String.Size(v.str)
	return size + slicev1qt0ZcfFi3cECΔrX3nfjAΞΞ.Size(v.sl)
}

func (s fooMUS) Skip(r mus.Reader) (n int, err error) {
  // First field: direct assignment to n (optimization)
	n, err = varint.Int.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = slicev1qt0ZcfFi3cECΔrX3nfjAΞΞ.Skip(r)
	n += n1
	return
}
```
