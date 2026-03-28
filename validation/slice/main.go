package main

import (
	"errors"
	"fmt"

	com "github.com/mus-format/common-go"
	slops "github.com/mus-format/mus-go/options/slice"
	"github.com/mus-format/mus-go/ord"
)

// This example demonstrates how to use slice length and element validators.
func main() {
	// 1. Create a slice value and a valid serializer.
	var (
		value []string

		// The length validator returns an error if the slice has more than 100 elements.
		lenVl com.ValidatorFn[int] = func(length int) error {
			if length > 100 {
				return com.ErrTooLargeLength
			}
			return nil
		}
		// The element validator returns an error if any element is empty.
		elemVl com.ValidatorFn[string] = func(elem string) error {
			if elem == "" {
				return errors.New("empty element")
			}
			return nil
		}

		// Each of the validators could be nil.
		ser = ord.NewValidSliceSer(
			ord.String,
			slops.WithLenValidator[string](lenVl),
			slops.WithElemValidator(elemVl),
		)
	)

	// Fill the slice to trigger the length validator.
	for i := 0; i < 101; i++ {
		value = append(value, "hello")
	}

	// 2. Calculate the required size.
	var (
		size = ser.Size(value)
		bs   = make([]byte, size)
	)

	// 3. Marshal the slice into the byte slice.
	n := ser.Marshal(value, bs)
	fmt.Printf("Marshal %d bytes\n", n)

	// 4. Unmarshal back into a new slice.
	// Unmarshalling stops immediately when any validator returns an error.
	// In this case, we expect a length validation error.
	value1, n, err := ser.Unmarshal(bs)
	if err != nil {
		fmt.Printf("Unmarshal failed as expected: %v\n", err)
	} else {
		fmt.Printf("Unmarshal %d bytes, Slice: %+v\n", n, value1)
	}
}
