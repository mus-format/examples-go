package main

import (
	assert "github.com/ymz-ncnk/assert/panic"
)

// This example demonstrates data versioning.
func main() {
	// 1. Marshal version V1.
	fooV1 := FooV1{num: 10}
	bs := make([]byte, FooV1DTS.Size(fooV1))
	FooV1DTS.Marshal(fooV1, bs)

	// 2. Unmarshal the current version (migrating from V1).
	foo, _, err := FooMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.Equal(foo, Foo{str: "10"})
}
