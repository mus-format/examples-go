# MUS Go Examples

The following examples demonstrate the **MUS serialization framework** (serving as a primary guide for both [mus](https://github.com/mus-format/mus-go) and [mus-stream](https://github.com/mus-format/mus-stream-go) libraries).

## Code Generation

- [mus-gen](mus-gen): Standard code generator.
- [mus-skill](mus-skill): AI-driven code generation.

## Manual Serialization

- [types](types): Handling of arrays, slices, maps, and structs.
- [validation](validation): Data validation during unmarshalling.
- [out_of_order](out_of_order): Deserializing values in reverse order.

## Typed Serialization & Multi-format

- [typed](typed): Versioning, interfaces, DTM, etc.
- [protobuf](protobuf): Implementing Protobuf-style encoding.

## Advanced Features

- [unsafe](unsafe): High-performance serialization.
- [pm](pm): Serializing linked lists and cyclic graphs (Pointer Mapping).
- [marshal_func](marshal_func): Using the `mus.Marshal` function.
