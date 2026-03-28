# Serialization Modes

The code could be generated in 3 built-in modes: 
- [Safe Mode](./mode_safe.md) (default)
- [Unsafe Mode](./mode_unsafe.md)
- [Not Unsafe Mode](./mode_not_unsafe.md)

And one custom mode:
- [Custom Mode](./mode_custom.md)

Where keys:
- **`int`**: refers to `int`, `int8`, `int16`, `int32`, `int64` types.
- **`uint`**: refers to `uint`, `uint8`, `uint16`, `uint32`, `uint64` types.
- **`float`**: refers to `float32`, `float64` types.
- **`byte`**: refers to `byte` type.
- **`bool`**: refers to `bool` type.
- **`string`**: refers to `string` type.
- **`byte_slice`**: refers to `[]byte` type.
- **`time`**: refers to `time.Time` type.
- **`time_ser`**: refers to `time.Time` serializer name.

Values typically refer to a package name (`varint`, `raw`, `unsafe`, `ord`), 
except for `time_ser` which specifies the serializer name within the package.

## Precedence

The generator resolves serializers in the following order of priority:

1. **Explicit Hints**: `// mus:` hints (e.g., `// mus:numEnc = raw`) in the source code.
2. **Mode Settings**: Settings defined by the selected built-in or custom mode.
