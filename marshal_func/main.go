package main

import (
	"fmt"

	"github.com/mus-format/ext-go"
)

func main() {
	// Both Foo and Bar types implement ext.MarshallerMUS interface.

	bs := ext.MarshalMUS(Foo{num: 10}) // ext.MarshalMUS can be used with Foo ...
	fmt.Println(bs)

	bs = ext.MarshalMUS(Bar{str: "10"}) // ... and with Bar.
	fmt.Println(bs)
}
