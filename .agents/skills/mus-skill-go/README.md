# mus-skill-go: AI-Driven MUS Serializer Generation

This repository provides a specialized AI skill to generate high-performance MUS 
serializers for Go.

## Contents

- [mus-skill-go: AI-Driven MUS Serializer Generation](#mus-skill-go-ai-driven-mus-serializer-generation)
  - [Contents](#contents)
  - [Setup](#setup)
  - [Supported Types](#supported-types)
  - [Usage](#usage)
  - [What to Expect](#what-to-expect)
  - [User Hints](#user-hints)
    - [Serializer Name](#serializer-name)
    - [Serializer Path](#serializer-path)
    - [Ignore Field](#ignore-field)
    - [Number Encoding](#number-encoding)
    - [Interface Serializer](#interface-serializer)
    - [Validation](#validation)
  - [Serialization Modes](#serialization-modes)
      - [.mus Configuration File](#mus-configuration-file)
        - [Hierarchy and Search Logic](#hierarchy-and-search-logic)
        - [Supported Configuration Values](#supported-configuration-values)
  - [Best Practices](#best-practices)

## Setup

Clone this repository into your project's agent directory (e.g. `.agent` or 
`.agents`):

```bash
git clone https://github.com/mus-format/mus-skill-go .agent/mus-skill-go
```

## Supported Types

You can generate serializers for the defined types only, like:

```go
type Foo struct {
    num int
}

type Bar int
```

## Usage

Ask the AI agent to generate MUS serializers. For example:

```text
Generate MUS serializers for the types found in the <file_name>.go file.
```

## What to Expect

The AI agent will generate the following files in your target and related packages:

- **`mus.ai.gen.go`**: Contains all generated serializers, including 
  implementation for defined types, structs, and interfaces.
- **`mus.ai.gen_test.go`**: Contains comprehensive unit tests for your serializers, 
  including validation logic.

After generation, you should always verify the output and run:

```bash
go test ./...
```

## User Hints

To customize the generation process use hints.

### Serializer Name

```go
// mus:name = CustomFoo
type Foo string

// The generated serializer will be named CustomFooMUS instead of FooMUS.
```

### Serializer Path

```go
// mus:path = github.com/user/repo/package
type Foo string

// path hint specifies the location of the type serializer.
```

### Ignore Field

```go
type Foo struct {
	num int
    // mus:ignore = true
    str string
}

// Ignored field will be skipped from the serialization process.
```

### Number Encoding

```go
// mus:numEnc = raw
type Foo int

type Bar struct {
  // mus:numEnc = raw
  num int
}

// Raw package will be used instead of a default varint.
```

### Interface Serializer

```go
// mus:impls = Foo, Bar
type MyInterface interface { ... }
```

You should define DTMs for ALL implementation types, or for none of them (in
this case DTMs will be generated automatically).

It's recommended to define DTMs yourself, because it's more flexible and 
maintainable.

### Validation

```go
// mus:vl = ValidateFoo
type Foo int

type Bar struct {
  // mus:vl = ValidateNum
  num int
}

type Zoo struct {
  // mus:lenVl = ValidateLength
  // mus:vl = ValidateStr
  str string

  // mus:elemVl = ValidateElement
  arr [3]int
    
  // mus:lenVl = ValidateLength
  bsl []byte
    
  // mus:lenVl = ValidateLength
  // mus:elemVl = ValidateElement
  sl []int

  // mus:lenVl = ValidateLength
  // mus:keyVl = ValidateKey
  // mus:elemVl = ValidateValue
  mp map[string]int
}
```

Validation function should be defined like:

```go
func Validate(T) error { ... }
```

Where `T` is the type being validated.

## Serialization Modes

There are 4 built-in serialization modes that control how the AI generates code 
(e.g., whether to use `unsafe` or not).

1. **Safe (Default)**: Optimized for safety. It does NOT use the `unsafe` 
   package. Numbers use `varint` encoding by default.
2. **Unsafe**: Optimized for maximum performance. It uses the `unsafe` package 
   for all types, including `string` and `byte slice`.
3. **Not Unsafe**: A middle ground. It uses the `unsafe` package for all types 
   EXCEPT `string`. This is often the fastest mode without the side effects 
   of unsafe string conversions.
4. **Custom**: Allows you to define your own rules via a `.mus` file.

So you can ask the AI agent to generate code, for example, in "unsafe" mode.

#### .mus Configuration File

For the **Custom** mode (or to override defaults in any mode), you can place a 
`.mus` file in your project. For example:

```yaml
# .mus
mode: custom
int: varint
uint: varint
float: unsafe
byte: raw
bool: ord
string: ord
byte_slice: ord
time: raw
time_ser: TimeUnixUTC
```

##### Hierarchy and Search Logic

The generator uses "ascending search" logic to find the appropriate `.mus` 
configuration:
1. It looks in the current directory.
2. If not found, it moves up to the parent directory until it reaches the 
   project root (where `go.mod` or `.git` is located).
3. **Overriding**: Specific `.mus` files in subdirectories override more general 
   ones found further up the tree.

##### Supported Configuration Values

| Key              | Values                                | Description                                             |
| :--------------- | :------------------------------------ | :------------------------------------------------------ |
| **`mode`**       | `custom`                              | Mode name.                                              |
| **`int`**        | `varint`, `raw`, `unsafe`             | Default package for all integer types.                  |
| **`uint`**       | `varint`, `raw`, `unsafe`             | Default package for all unsigned integer types.         |
| **`float`**      | `varint`, `raw`, `unsafe`             | Default package for all float types.                    |
| **`bool`**       | `ord`, `unsafe`                       | Default package for `bool`.                             |
| **`string`**     | `ord`, `unsafe`                       | Default package for `string`.                           |
| **`byte`**       | `varint`, `raw`, `unsafe`             | Default package for `byte`.                             |
| **`byte_slice`** | `ord`, `unsafe`                       | Default package for `[]byte`.                           |
| **`time`**       | `raw`, `unsafe`                       | Default package for `time.Time`.                        |
| **`time_ser`**   | `TimeUnix`, `TimeUnixMilli`,          | Specific `time.Time` serializer from `raw` or `unsafe`. |
|                  | `TimeUnixMicro`, `TimeUnixNano`,      |                                                         |
|                  | `TimeUnixUTC`, `TimeUnixMilliUTC`,    |                                                         |
|                  | `TimeUnixMicroUTC`, `TimeUnixNanoUTC` |                                                         |

## Best Practices

- **Default to Safe**: Use the default `safe` mode unless you have a confirmed 
  performance bottleneck.
- **Manual DTMs**: For interfaces, define DTM constants yourself for better 
  stability.
- **Validator Signatures**: Ensure your validator functions match the 
  `func(T) error` signature.
