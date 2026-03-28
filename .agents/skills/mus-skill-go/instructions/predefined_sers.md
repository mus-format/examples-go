# Predefined Serializers

There are a set of predefined serializers that can be used to build more complex 
serializers. Let's see on packages.

## varint

Contains Varint serializers for the following number types:

- `uint8`: varint.Uint8
- `uint16`: varint.Uint16
- `uint32`: varint.Uint32
- `uint64`: varint.Uint64
- `uint`: varint.Uint
- `int`: varint.Int
- `int16`: varint.Int16
- `int32`: varint.Int32
- `int64`: varint.Int64
- `int`: varint.Int
- `float32`: varint.Float32
- `float64`: varint.Float64

## raw

Contains Raw serializers for the following number types:

- `uint8`: raw.Uint8
- `uint16`: raw.Uint16
- `uint32`: raw.Uint32
- `uint64`: raw.Uint64
- `uint`: raw.Uint
- `int`: raw.Int
- `int16`: raw.Int16
- `int32`: raw.Int32
- `int64`: raw.Int64
- `int`: raw.Int
- `float32`: raw.Float32
- `float64`: raw.Float64
- `time.Time`: raw.TimeUnix, raw.TimeUnixMilli, raw.TimeUnixMicro, 
  raw.TimeUnixNano, raw.TimeUnixUTC, raw.TimeUnixMilliUTC, raw.TimeUnixMicroUTC,
	raw.TimeUnixNanoUTC

## ord

Contains serializers for the following types:

- `bool`: ord.Bool
- `[]byte`: ord.ByteSlice
- `string`: ord.String

## unsafe

Contains serializers for the following types:

- `uint8`: unsafe.Uint8
- `uint16`: unsafe.Uint16
- `uint32`: unsafe.Uint32
- `uint64`: unsafe.Uint64
- `uint`: unsafe.Uint
- `int`: unsafe.Int
- `int16`: unsafe.Int16
- `int32`: unsafe.Int32
- `int64`: unsafe.Int64
- `int`: unsafe.Int
- `float32`: unsafe.Float32
- `float64`: unsafe.Float64
- `bool`: unsafe.Bool
- `[]byte`: unsafe.ByteSlice
- `string`: unsafe.String
- `time.Time`: unsafe.TimeUnix, unsafe.TimeUnixMilli, unsafe.TimeUnixMicro,
  unsafe.TimeUnixNano, unsafe.TimeUnixUTC, unsafe.TimeUnixMilliUTC, 
	unsafe.TimeUnixMicroUTC, unsafe.TimeUnixNanoUTC
