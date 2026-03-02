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
		value = map[int]string{1: "hello", 2: "world"}

		// The length validator returns an error if the map has more than 3 elements.
		lenVl com.ValidatorFn[int] = func(length int) error {
			if length > 3 {
				return com.ErrTooLargeLength
			}
			return nil
		}
		// The key validator returns an error if a key equals 1.
		keyVl com.ValidatorFn[int] = func(key int) error {
			if key == 1 {
				return errors.New("bad key")
			}
			return nil
		}
		// The value validator returns an error if a value equals "hello".
		valueVl com.ValidatorFn[string] = func(val string) error {
			if val == "hello" {
				return errors.New("bad value")
			}
			return nil
		}

		// Each of the validators could be nil.
		ser = ord.NewValidMapSer(
			varint.Int,
			ord.String,
			mapops.WithLenValidator[int, string](lenVl),
			mapops.WithKeyValidator[int, string](keyVl),
			mapops.WithValueValidator[int](valueVl),
		)

		// To specify the length serializer use:
		// ser = ord.NewValidMapSer[int, string](varint.Int, ord.String,
		//    mapops.WithLenSer[int, string](lenSer), ...)
	)

	// 2. Calculate the required size.
	var (
		size = ser.Size(value)
		bs   = make([]byte, size)
	)

	// 3. Marshal the map into the byte slice.
	n := ser.Marshal(value, bs)
	fmt.Printf("Marshal %d bytes\n", n)

	// 4. Unmarshal back into a new map.
	// Unmarshalling stops immediately when any validator returns an error.
	// In this case, we expect a key validation error.
	value1, n, err := ser.Unmarshal(bs)
	if err != nil {
		fmt.Printf("Unmarshal failed as expected: %v\n", err)
	} else {
		fmt.Printf("Unmarshal %d bytes, Map: %+v\n", n, value1)
	}
}
