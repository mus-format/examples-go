package main

import "fmt"

// This example demonstrates how to serialize a struct.
func main() {
	// 1. Create a Foo instance.
	v := Foo{
		str: "Hello, MUS!",
		sl:  []int{1, 2, 3},
	}

	// 2. Calculate the required size.
	// This ensures we allocate exactly the amount of memory needed.
	var (
		size = FooMUS.Size(v)
		bs   = make([]byte, size)
	)

	// 3. Marshal the struct into the byte slice.
	n := FooMUS.Marshal(v, bs)
	fmt.Printf("Marshaled %+v → %d bytes: %x\n", v, n, bs)

	// 4. Unmarshal back into a new struct.
	v2, n, err := FooMUS.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshaled back: %+v\n", v2)

	// 5. Demonstrate Skip.
	// This returns the number of bytes occupied by the Foo struct without
	// decoding.
	n, err = FooMUS.Skip(bs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Skipped %d bytes\n", n)
}
