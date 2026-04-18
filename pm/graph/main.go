package main

import (
	"fmt"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
	assert "github.com/ymz-ncnk/assert/panic"
)

func init() {
	assert.On = true
}

// This example demonstrates how to use the pm package to serialize a cyclic
// graph.
func main() {
	var (
		v        = CyclicGraph()
		graphMUS = MakeGraphMUS(varint.Int, ord.String)
	)

	// 1. Marshal graph.
	bs := make([]byte, graphMUS.Size(v))
	graphMUS.Marshal(v, bs)

	// 2. Unmarshal graph.
	av, _, err := graphMUS.Unmarshal(bs)
	assert.EqualError(err, nil)
	assert.EqualDeep(v, av)

	fmt.Printf("Marshaled %v → %d bytes: %x\n", v, len(bs), bs)
	fmt.Printf("Unmarshaled back: %v\n", av)
}

func CyclicGraph() (g Graph[int, string]) {
	g = NewGraph[int, string]()
	g.AddVertex(1, "one")
	g.AddVertex(2, "two")
	g.AddVertex(3, "three")

	g.AddEdge(1, 2, 10)
	g.AddEdge(2, 3, 20)
	g.AddEdge(3, 1, 30)
	return
}
