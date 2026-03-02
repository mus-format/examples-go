package main

import (
	"fmt"
	"math/rand"

	com "github.com/mus-format/common-go"
	dts "github.com/mus-format/dts-go"
)

// This example demonstrates how to use DTM (Data Type Metadata) to handle
// different data types.
func main() {
	// 1. Unmarshal DTM (Data Type Metadata) to identify the data type.
	bs := randomData()
	dtm, n, err := dts.DTMSer.Unmarshal(bs)
	if err != nil {
		panic(err)
	}

	// 2. Unmarshal the actual data based on the DTM.
	switch dtm {
	case FooDTM:
		foo, _, err := FooDTS.UnmarshalData(bs[n:])
		if err != nil {
			panic(err)
		}
		// process foo ...
		fmt.Println(foo)
	case BarDTM:
		bar, _, err := BarDTS.UnmarshalData(bs[n:])
		if err != nil {
			panic(err)
		}
		// process bar ...
		fmt.Println(bar)
	default:
		panic(fmt.Sprintf("unexpected %v DTM", dtm))
	}
}

func randomData() (bs []byte) {
	// 1. Generate a random DTM (Data Type Metadata).
	dtm := com.DTM(rand.Intn(2) + 1)
	switch dtm {
	// 2. Marshal the corresponding data type.
	case FooDTM:
		foo := Foo{num: 5}
		bs = make([]byte, FooDTS.Size(foo))
		FooDTS.Marshal(foo, bs)
	case BarDTM:
		bar := Bar{str: "hello world"}
		bs = make([]byte, BarDTS.Size(bar))
		BarDTS.Marshal(bar, bs)
	default:
		panic(fmt.Sprintf("unexpected %v DTM", dtm))
	}
	return
}
