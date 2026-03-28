package main

import (
	"fmt"

	assert "github.com/ymz-ncnk/assert/panic"
)

func init() {
	assert.On = true
}

// This example demonstrates how the pm package works.
func main() {
	SerializeTwoPtrs()
	SerializeThreePtrs()
}

func SerializeTwoPtrs() {
	var (
		twoPtrsMUS = MakeTwoPtrsSer()
		v          = NewTwoPtrs("the same pointer in two fields")
	)

	// 1. Marshal TwoPtrs.
	bs := make([]byte, twoPtrsMUS.Size(v))
	twoPtrsMUS.Marshal(v, bs)

	// 2. Unmarshal TwoPtrs.
	av, _, err := twoPtrsMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.Equal(av.ptr1, av.ptr2)

	fmt.Printf("Two pointers: %+v\n", av)
}

func SerializeThreePtrs() {
	var (
		// ThreePtrs structure serializer uses TwoPtrs serializer.
		threePtrsMUS = MakeThreePtrsSer()
		v            = NewThreePtrs("the same pointer in three fields")
	)

	// 1. Marshal ThreePtrs.
	bs := make([]byte, threePtrsMUS.Size(v))
	threePtrsMUS.Marshal(v, bs)

	// 2. Unmarshal ThreePtrs.
	av, _, err := threePtrsMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.Equal(av.ptr1, av.ptr2)
	assert.Equal(av.ptr1, av.ptr3)
	fmt.Printf("Three pointers: %+v\n", av)
}

func NewTwoPtrs(str string) TwoPtrs {
	ptr := &str
	return TwoPtrs{
		ptr1: ptr,
		ptr2: ptr,
	}
}

func NewThreePtrs(str string) ThreePtrs {
	ptr := &str
	return ThreePtrs{
		TwoPtrs: TwoPtrs{
			ptr1: ptr,
			ptr2: ptr,
		},
		ptr3: ptr,
	}
}
