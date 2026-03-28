package main

import (
	"errors"
	"fmt"

	com "github.com/mus-format/common-go"
	mapops "github.com/mus-format/mus-go/options/map"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

// This example demonstrates how to use map length, key and value validators.
func main() {
	// 1. Create a map value and a valid serializer.
	var (
		value = make(map[int]string)

		// The length validator returns an error if the map has more than 100 elements.
		lenVl com.ValidatorFn[int] = func(length int) error {
			if length > 100 {
				return com.ErrTooLargeLength
			}
			return nil
		}
		// The key validator returns an error if a key is negative.
		keyVl com.ValidatorFn[int] = func(key int) error {
			if key < 0 {
				return errors.New("negative key")
			}
			return nil
		}
		// The value validator returns an error if a value is empty.
		valueVl com.ValidatorFn[string] = func(val string) error {
			if val == "" {
				return errors.New("empty value")
			}
			return nil
		}

		// Each of the validators could be nil.
		mapMUS = ord.NewValidMapSer(
			varint.Int,
			ord.String,
			mapops.WithLenValidator[int, string](lenVl),
			mapops.WithKeyValidator[int, string](keyVl),
			mapops.WithValueValidator[int](valueVl),
		)
	)

	// Fill the map to trigger the length validator.
	for i := range 101 {
		value[i] = "hello"
	}

	// 2. Calculate the required size.
	var (
		size = mapMUS.Size(value)
		bs   = make([]byte, size)
	)

	// 3. Marshal the map into the byte slice.
	n := mapMUS.Marshal(value, bs)
	fmt.Printf("Marshal %d bytes\n", n)

	// 4. Unmarshal back into a new map.
	// Unmarshalling stops immediately when any validator returns an error.
	// In this case, we expect a length validation error.
	value1, n, err := mapMUS.Unmarshal(bs)
	if err != nil {
		fmt.Printf("Unmarshal failed as expected: %v\n", err)
	} else {
		fmt.Printf("Unmarshal %d bytes, Map: %+v\n", n, value1)
	}
}
