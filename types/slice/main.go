package main

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
	// slops "github.com/mus-format/mus-go/options/slice"
)

// This example demonstrates how to serialize a slice.
func main() {
	// 1. Initialize the serializer
	// varint.Int is the serializer of the slice's elements.
	ser := ord.NewSliceSer(varint.Int)
	// To specify length serializer use:
	// ser = ord.NewSliceSer[int](varint.Int, slops.WithLenSer(lenSer))

	// 2. Marshal
	var (
		sl   = []int{1, 2, 3}
		size = ser.Size(sl)
		bs   = make([]byte, size)
	)
	n := ser.Marshal(sl, bs)
	fmt.Printf("Marshaled %d bytes\n", n)

	// 3. Unmarshal
	sl, n, err := ser.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshaled %d bytes, slice: %v\n", n, sl)
}
