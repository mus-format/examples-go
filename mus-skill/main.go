package main

import (
	"errors"
	"fmt"

	assert "github.com/ymz-ncnk/assert/panic"
)

// mus:name = CustomFoo
type Foo struct {
	// mus:numEnc = raw
	Num int
	// mus:lenVl = ValidateLength
	Str string
}

func ValidateLength(l int) error {
	if l > 100 {
		return errors.New("string too long")
	}
	return nil
}

func main() {
	foo := Foo{Num: 1, Str: "hello"}
	bs := make([]byte, CustomFooMUS.Size(foo))
	CustomFooMUS.Marshal(foo, bs)

	afoo, _, err := CustomFooMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(foo, afoo)

	fmt.Printf("Marshaled %+v → %d bytes: %x\n", foo, len(bs), bs)
	fmt.Printf("Unmarshaled back: %+v\n", afoo)
}
