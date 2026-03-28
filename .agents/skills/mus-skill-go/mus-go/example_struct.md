# Struct Example

```go
type Foo struct {
	// mus:vl = ValidateNum
	num int 
	str string
	// mus:lenEnc = raw
	sl  []int 
}

var sliceMq6O0A1hiHDBaP9SIUTMZQΞΞ = ord.NewSliceSer[int](varint.Int)

var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
	n = varint.Int.Marshal(v.num, bs)
	n += ord.String.Marshal(v.str, bs[n:])
	return n + sliceMq6O0A1hiHDBaP9SIUTMZQΞΞ.Marshal(v.sl, bs[n:])
}

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
  // First field: direct assignment to n (optimization)
	v.num, n, err = varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	err = ValidateNum(v.num)
	if err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.sl, n1, err = sliceMq6O0A1hiHDBaP9SIUTMZQΞΞ.Unmarshal(bs[n:])
	n += n1
	return
}

func (s fooMUS) Size(v Foo) (size int) {
	size = varint.Int.Size(v.num)
	size += ord.String.Size(v.str)
	return size + sliceMq6O0A1hiHDBaP9SIUTMZQΞΞ.Size(v.sl)
}

func (s fooMUS) Skip(bs []byte) (n int, err error) {
  // First field: direct assignment to n (optimization)
	n, err = varint.Int.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = sliceMq6O0A1hiHDBaP9SIUTMZQΞΞ.Skip(bs[n:])
	n += n1
	return
}
```
