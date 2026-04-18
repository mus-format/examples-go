package main

import (
	"fmt"
	"math/rand"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go/typed"
)

// This example demonstrates how to use DTM (Data Type Metadata) to handle
// different data types.
func main() {
	bs, v := randomData()
	fmt.Printf("Marshaled %T %+v → %d bytes: %x\n", v, v, len(bs), bs)

	// 1. Unmarshal DTM (Data Type Metadata) to identify the data type.
	dtm, n, err := typed.DTMSer.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	// 2. Unmarshal the actual data based on the DTM.
	switch dtm {
	case FooDTM:
		foo, _, err := FooTypedMUS.UnmarshalData(bs[n:])
		if err != nil {
			panic(err)
		}
		// process foo ...
		fmt.Printf("Unmarshaled Foo back: %+v\n", foo)
	case BarDTM:
		bar, _, err := BarTypedMUS.UnmarshalData(bs[n:])
		if err != nil {
			panic(err)
		}
		// process bar ...
		fmt.Printf("Unmarshaled Bar back: %+v\n", bar)
	default:
		panic(fmt.Sprintf("unexpected %v DTM", dtm))
	}
}

func randomData() (bs []byte, v any) {
	// 1. Generate a random DTM (Data Type Metadata).
	dtm := com.DTM(rand.Intn(2) + 1)
	switch dtm {
	// 2. Marshal the corresponding typed data.
	case FooDTM:
		foo := Foo{num: 5}
		v = foo
		bs = make([]byte, FooTypedMUS.Size(foo))
		FooTypedMUS.Marshal(foo, bs)
	case BarDTM:
		bar := Bar{str: "hello world"}
		v = bar
		bs = make([]byte, BarTypedMUS.Size(bar))
		BarTypedMUS.Marshal(bar, bs)
	default:
		panic(fmt.Sprintf("unexpected %v DTM", dtm))
	}
	return
}
