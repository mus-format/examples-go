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
		arr = [3]int{1, 2, 3}
		ser = unsafe.NewArraySer[[3]int](varint.Int)
	)

	// 2. Calculate the required size.
	// This ensures we allocate exactly the amount of memory needed.
	var (
		size = ser.Size(arr)
		bs   = make([]byte, size)
	)

	// 3. Marshal the array into the byte slice.
	n := ser.Marshal(arr, bs)
	fmt.Printf("Marshal %d bytes\n", n)

	// 4. Unmarshal back into a new array.
	arr1, n, err := ser.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshal %d bytes, Array: %+v\n", n, arr1)

	// 5. Demonstrate Skip.
	// This returns the number of bytes occupied by the array without decoding.
	n, err = ser.Skip(bs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Skip %d bytes\n", n)
}
