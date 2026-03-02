package main

import (
	"fmt"

	"github.com/mus-format/mus-go/varint"
)

// This example demonstrates how to deserialize values in reverse order.
func main() {
	// 1. Encode three numbers in sequence: 5, 10, 15.
	var (
		size = varint.Int.Size(5) + varint.Int.Size(10) + varint.Int.Size(15)
		bs   = make([]byte, size)
	)
	n := varint.Int.Marshal(5, bs)
	n += varint.Int.Marshal(10, bs[n:])
	varint.Int.Marshal(15, bs[n:])

	// 2. Access the data out of order (Reverse).
	// We use Skip to find the offsets without decoding the actual values.
	offset1, _ := varint.Int.Skip(bs)           // Skip 5
	offset2, _ := varint.Int.Skip(bs[offset1:]) // Skip 10

	// Get the third number (15)
	num, _, _ := varint.Int.Unmarshal(bs[offset1+offset2:])
	fmt.Println(num) // Output: 15

	// Get the second number (10)
	num, _, _ = varint.Int.Unmarshal(bs[offset1:])
	fmt.Println(num) // Output: 10

	// Get the first number (5)
	num, _, _ = varint.Int.Unmarshal(bs)
	fmt.Println(num) // Output: 5
}
