package main

import (
	"fmt"
)

// Interface to Marshal/Unmarshal.
type Instruction interface {
	Do()
}

// Copy implements Instruction and Marshaller interfaces.
type Copy struct {
	start int
	end   int
}

func (c Copy) Do() {
	fmt.Printf("copy from %v to %v\n", c.start, c.end)
}

func (c Copy) MarshalTypedMUS(bs []byte) (n int) {
	return CopyTypedMUS.Marshal(c, bs)
}

func (c Copy) SizeTypedMUS() int {
	return CopyTypedMUS.Size(c)
}

// Insert implements Instruction and Marshaller interfaces.
type Insert struct {
	str string
}

func (i Insert) Do() {
	fmt.Printf("insert '%v'\n", i.str)
}

func (i Insert) MarshalTypedMUS(bs []byte) (n int) {
	return InsertTypedMUS.Marshal(i, bs)
}

func (i Insert) SizeTypedMUS() int {
	return InsertTypedMUS.Size(i)
}
