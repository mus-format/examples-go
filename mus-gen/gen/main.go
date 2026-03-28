package main

import (
	"log"
	"os"
	"reflect"

	"github.com/mus-format/examples-go/mus-gen/pkg"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genopts "github.com/mus-format/mus-gen-go/options/gen"
)

func main() {
	g, err := musgen.NewGenerator(
		genopts.WithPkgPath("github.com/mus-format/examples-go/mus-gen/pkg"),
		// genopts.WithPackage("bar"), // Can be used to specify the package name for
		// the generated file.
	)
	if err != nil {
		panic(err)
	}
	err = g.AddDefinedType(reflect.TypeFor[pkg.Int]())
	if err != nil {
		panic(err)
	}
	err = g.AddStruct(reflect.TypeFor[pkg.Foo[pkg.Int]]())
	if err != nil {
		panic(err)
	}
	bs, err := g.Generate()
	if err != nil {
		// In case of an error (e.g., if you forget to specify an import path using
		// genopts.WithImport), the generated code can be inspected for additional
		// details.
		log.Println(err)
	}
	err = os.WriteFile("./pkg/mus.gen.go", bs, 0644)
	if err != nil {
		panic(err)
	}
}
