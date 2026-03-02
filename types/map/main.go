package main

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
	// mapops "github.com/mus-format/mus-go/options/map"
)

// This example demonstrates how to serialize a map.
func main() {
	// 1. Initialize the serializer
	// varint.Int specifies the serializer for the map’s keys, and ord.String -
	// the serializer for the map’s values.
	ser := ord.NewMapSer(varint.Int, ord.String)

	// To specify length serializer use:
	// ser = ord.NewMapSer[int, string](varint.Int, ord.String, mapops.WithLenSer(lenSer))

	// 2. Marshal
	var (
		m    = map[int]string{1: "one", 2: "two", 3: "three"}
		size = ser.Size(m)
		bs   = make([]byte, size)
	)
	n := ser.Marshal(m, bs)
	fmt.Printf("Marshaled %d bytes\n", n)

	// 3. Unmarshal
	m, n, err := ser.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshaled %d bytes, map: %v\n", n, m)
}
