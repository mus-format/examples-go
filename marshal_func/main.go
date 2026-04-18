package main

import (
	"fmt"

	"github.com/mus-format/mus-go"
)

// This example demonstrates how to use the mus.Marshal function.
func main() {
	// Marshal Foo (numeric field) and print the result.
	bs := mus.Marshal(Foo{num: 10})
	fmt.Printf("Marshaled Foo: %v\n", bs)

	// Marshal Bar (string field) and print the result.
	bs = mus.Marshal(Bar{str: "10"})
	fmt.Printf("Marshaled Bar: %v\n", bs)
}
