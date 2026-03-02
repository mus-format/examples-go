package main

import (
	"fmt"

	"github.com/mus-format/ext-go"
)

// This example demonstrates how to use the MarshalMUS function.
func main() {
	// 1. Marshal Foo.
	bs := ext.MarshalMUS(Foo{num: 10})
	fmt.Println(bs)

	// 2. Marshal Bar.
	bs = ext.MarshalMUS(Bar{str: "10"})
	fmt.Println(bs)
}
