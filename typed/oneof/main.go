package main

import (
	"fmt"

	assert "github.com/ymz-ncnk/assert/panic"
)

func init() {
	assert.On = true
}

// This example demonstrates how to serialize and deserialize Go interfaces.
func main() {
	var (
		bs  []byte
		in  Instruction // Interface.
		err error
	)

	// 1. Marshal/Unmarshal Copy instruction.
	copy := Copy{start: 10, end: 20}
	bs = make([]byte, InstructionMUS.Size(copy))
	InstructionMUS.Marshal(copy, bs)
	fmt.Printf("Marshaled Copy %+v → %d bytes: %x\n", copy, len(bs), bs)

	in, _, err = InstructionMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(in, copy)
	fmt.Printf("Unmarshaled Instruction back: %+v\n\n", in)

	// 2. Marshal/Unmarshal Insert instruction.
	insert := Insert{str: "hello world"}
	bs = make([]byte, InstructionMUS.Size(insert))
	InstructionMUS.Marshal(insert, bs)
	fmt.Printf("Marshaled Insert %+v → %d bytes: %x\n", insert, len(bs), bs)

	in, _, err = InstructionMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(in, insert)
	fmt.Printf("Unmarshaled Instruction back: %+v\n", in)
}
