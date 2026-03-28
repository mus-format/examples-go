package main

import (
	"fmt"
	"strings"

	com "github.com/mus-format/common-go"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/ord"
)

// This example demonstrates how to use string length validator.
func main() {
	// 1. Create a string value and a valid serializer.
	var (
		value = strings.Repeat("a", 101)

		// The validator will return an error if the string length exceeds 100.
		lenVl com.ValidatorFn[int] = func(length int) error {
			if length > 100 {
				return com.ErrTooLargeLength
			}
			return nil
		}
		ser = ord.NewValidStringSer(strops.WithLenValidator(lenVl))
	)

	// 2. Calculate the required size.
	var (
		size = ser.Size(value)
		bs   = make([]byte, size)
	)

	// 3. Marshal the string into the byte slice.
	n := ser.Marshal(value, bs)
	fmt.Printf("Marshal %d bytes\n", n)

	// 4. Unmarshal back into a new string.
	// Unmarshalling stops immediately when a validator returns an error.
	// In this case, we expect a length validation error.
	value1, n, err := ser.Unmarshal(bs)
	if err != nil {
		fmt.Printf("Unmarshal failed as expected: %v\n", err)
	} else {
		fmt.Printf("Unmarshal %d bytes, String: %s\n", n, value1)
	}
}
