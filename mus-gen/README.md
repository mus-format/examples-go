# mus-gen Example

This example demonstrates how to use the [mus-gen-go](https://github.com/mus-format/mus-gen-go) 
tool for code generation.

## How it Works

1. **Define Types**: Go types are defined in the `pkg` directory (e.g., `pkg/types.go`).
2. **Generator Script**: A generator script is created in `gen/main.go`. It 
   configures the `mus-gen` generator with the package path and types.
3. **Generate Code**: The `go:generate` directive in `main.go` runs the generator 
   script, which produces `pkg/mus.gen.go`.
4. **Usage**: The generated serializers are used in `main.go`.

## Running the Example

From the `mus-gen` directory:

```bash
# Generate the serializers
go generate ./...

# Run the example
go run .
```
