//go:generate go run gen/main.go
package main

import (
	"fmt"

	"github.com/mus-format/examples-go/mus-gen/pkg"
	assert "github.com/ymz-ncnk/assert/panic"
)

// Enable assertions.
func init() {
	assert.On = true
}

func main() {
	var (
		// Sample value.
		foo = pkg.Foo[pkg.Int]{
			S: "hello world",
			T: pkg.Int(5),
		}
		// Allocate bs.
		size = pkg.FooMUS.Size(foo)
		bs   = make([]byte, size)
	)
	// Serialize / deserialize.
	pkg.FooMUS.Marshal(foo, bs)
	afoo, _, err := pkg.FooMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(foo, afoo)

	fmt.Printf("Marshaled %+v → %d bytes: %x\n", foo, len(bs), bs)
	fmt.Printf("Unmarshaled back: %+v\n", afoo)
}
