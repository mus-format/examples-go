package main

import (
	"fmt"

	"github.com/mus-format/mus-go"
)

// This example demonstrates how to use the mus.Marshal function.
func main() {
	// 1. Marshal Foo.
	bs := mus.Marshal(Foo{num: 10})
	fmt.Printf("Marshalled Foo: %v\n", bs)

	// 2. Marshal Bar.
	bs = mus.Marshal(Bar{str: "10"})
	fmt.Printf("Marshalled Bar: %v\n", bs)
}
