---
name: mus-skill-go
description: Generate MUS serializers for Go types.
---

# MUS Generation Skill

This skill provides modular instructions for generating MUS serializers for Go 
types. It supports both `mus-go` (`[]byte` based) and `mus-stream-go` 
(`io.Writer`/`io.Reader` based) flavors.

## Entry Point

Start with the [Instructions](./instructions/instructions.md) to understand the 
high-level generation workflow.

## Flavor Specifics

Based on the target flavor, refer to the following:

### mus-go (Buffer-based)
- **[Template](./mus-go/template.md)**: Serializer structure.
- **[Imports](./mus-go/imports.md)**: Required packages and aliases.
- **[Struct Example](./mus-go/example_struct.md)**
- **[Defined Type Example](./mus-go/example_defined_type.md)**
- **[Embedding Example](./mus-go/example_struct_embedding.md)**

### mus-stream-go (Stream-based)
- **[Template](./mus-stream-go/template.md)**: Serializer structure.
- **[Imports](./mus-stream-go/imports.md)**: Required packages and aliases.
- **[Struct Example](./mus-stream-go/example_struct.md)**
- **[Defined Type Example](./mus-stream-go/example_defined_type.md)**
- **[Embedding Example](./mus-stream-go/example_struct_embedding.md)**

## Utility Scripts

- **[Hash Generator](./scripts/hash_gen.py)**: Python script for generating 
  anonymous serializer names. Usage: `python3 scripts/hash_gen.py <prefix> <data>`.
