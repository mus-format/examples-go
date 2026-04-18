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
	SerializeEmbeddedPtrs()
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

	fmt.Printf("Marshaled %+v → %d bytes: %x\n", v, len(bs), bs)
	fmt.Printf("Unmarshaled back: %+v\n", av)
	fmt.Printf("Pointer identity: %p == %p\n\n", av.ptr1, av.ptr2)
}

func SerializeEmbeddedPtrs() {
	var (
		// EmbeddedPtrs structure serializer uses TwoPtrs serializer.
		embeddedPtrsMUS = MakeEmbeddedPtrsSer()
		v               = NewEmbeddedPtrs("the same pointer in composite struct")
	)

	// 1. Marshal EmbeddedPtrs.
	bs := make([]byte, embeddedPtrsMUS.Size(v))
	embeddedPtrsMUS.Marshal(v, bs)

	// 2. Unmarshal EmbeddedPtrs.
	av, _, err := embeddedPtrsMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.Equal(av.ptr1, av.ptr2)
	assert.Equal(av.ptr1, av.ptr3)
	fmt.Printf("Marshaled %+v → %d bytes: %x\n", v, len(bs), bs)
	fmt.Printf("Unmarshaled back: %+v\n", av)
	fmt.Printf("Pointer identity: %p == %p == %p\n", av.ptr1, av.ptr2, av.ptr3)
}

func NewTwoPtrs(str string) TwoPtrs {
	ptr := &str
	return TwoPtrs{
		ptr1: ptr,
		ptr2: ptr,
	}
}

func NewEmbeddedPtrs(str string) EmbeddedPtrs {
	ptr := &str
	return EmbeddedPtrs{
		TwoPtrs: TwoPtrs{
			ptr1: ptr,
			ptr2: ptr,
		},
		ptr3: ptr,
	}
}
