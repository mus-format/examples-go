//go:generate go run gen/main.go
package main

import (
	"fmt"

	"github.com/mus-format/examples-go/mus-gen/pkg"
	assert "github.com/ymz-ncnk/assert/panic"
)

func init() {
	assert.On = true
}

func main() {
	var (
		foo = pkg.Foo[pkg.Int]{
			S: "hello world",
			T: pkg.Int(5),
		}
		size = pkg.FooMUS.Size(foo)
		bs   = make([]byte, size)
	)
	pkg.FooMUS.Marshal(foo, bs)
	afoo, _, err := pkg.FooMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(foo, afoo)

	fmt.Printf("Generated serializer result: %+v\n", afoo)
}
