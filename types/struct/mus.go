package main

import (
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

var (
	// FooMUS is a serializer for the Foo struct.
	FooMUS = fooMUS{}

	// IntSliceMUS is an auxiliary serializer for the Foo's slice field.
	IntSliceMUS = ord.NewSliceSer(varint.Int)
)

// fooMUS implements the mus.Serializer interface.
type fooMUS struct{}

// Marshal fills bs with an encoded Foo value.
func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
	// 1. Marshal the first field.
	n = ord.String.Marshal(v.str, bs)
	// 2. Marshal the second field.
	return n + IntSliceMUS.Marshal(v.sl, bs[n:])
}

// Unmarshal parses an encoded Foo value from bs.
func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
	// 1. Unmarshal the first field.
	v.str, n, err = ord.String.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	// 2. Unmarshal the second field.
	v.sl, n1, err = IntSliceMUS.Unmarshal(bs[n:])
	// 3. Accumulate total used bytes.
	n += n1
	return
}

// Size calculates the total number of bytes required to serialize the Foo
// struct.
func (s fooMUS) Size(v Foo) (size int) {
	size += ord.String.Size(v.str)
	return size + IntSliceMUS.Size(v.sl)
}

// Skip returns the number of bytes occupied by the Foo object without fully
// decoding its fields.
func (s fooMUS) Skip(bs []byte) (n int, err error) {
	// 1. Skip the first field.
	n, err = ord.String.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	// 2. Skip the second field.
	n1, err = IntSliceMUS.Skip(bs[n:])
	n += n1
	return
}
