package main

import (
	"fmt"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
)

// This example demonstrates how to serialize an array.
func main() {
	// 1. Create an array instance and a serializer.
	var (
		arr    = [3]int{1, 2, 3}
		arrMUS = unsafe.NewArraySer[[3]int](varint.Int)
	)

	// 2. Calculate the required size.
	// This ensures we allocate exactly the amount of memory needed.
	var (
		size = arrMUS.Size(arr)
		bs   = make([]byte, size)
	)

	// 3. Marshal the array into the byte slice.
	n := arrMUS.Marshal(arr, bs)
	fmt.Printf("Marshaled %+v → %d bytes: %x\n", arr, n, bs)

	// 4. Unmarshal back into a new array.
	arr1, n, err := arrMUS.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshaled back: %+v\n", arr1)

	// 5. Demonstrate Skip.
	// This returns the number of bytes occupied by the array without decoding.
	n, err = arrMUS.Skip(bs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Skipped %d bytes\n", n)
}
