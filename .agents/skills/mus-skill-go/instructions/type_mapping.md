# Type Mapping

| Go Type   | MUS Type   | Default Serializer | Constructors               | Should create anonymous serializer? |
| --------- | ---------- | ------------------ | -------------------------- | ----------------------------------- |
| int       | int        | varint.Int         | -                          | -                                   |
| int8      | int8       | varint.Int8        | -                          | -                                   |
| int16     | int16      | varint.Int16       | -                          | -                                   |
| int32     | int32      | varint.Int32       | -                          | -                                   |
| int64     | int64      | varint.Int64       | -                          | -                                   |
| uint      | uint       | varint.Uint        | -                          | -                                   |
| uint8     | uint8      | raw.Uint8          | -                          | -                                   |
| uint16    | uint16     | varint.Uint16      | -                          | -                                   |
| uint32    | uint32     | varint.Uint32      | -                          | -                                   |
| uint64    | uint64     | varint.Uint64      | -                          | -                                   |
| byte      | uint8      | raw.Uint8          | -                          | -                                   |
| float32   | float32    | varint.Float32     | -                          | -                                   |
| float64   | float64    | varint.Float64     | -                          | -                                   |
| bool      | bool       | ord.Bool           | -                          | -                                   |
| string    | string     | ord.String         | ord.NewStringSer(),        | If user provides hints              |
| [N]T      | array      | -                  | unsafe.NewArraySer(),      | Always                              |
|           |            |                    | unsafe.NewValidArraySer()  |                                     |
| []T       | slice      | -                  | ord.NewSliceSer(),         | Always                              |
|           |            |                    | ord.NewValidSliceSer()     |                                     |
| []byte    | byte slice | ord.ByteSlice      | ord.NewByteSliceSer(),     | If user provides hints              |
|           |            |                    | ord.NewValidByteSliceSer() |                                     |
| map[K]V   | map        | -                  | ord.NewMapSer(),           | Always                              |
|           |            |                    | ord.NewValidMapSer()       |                                     |
| *T        | pointer    | -                  | ord.NewPtrSer(),           | Always                              |
| time.Time | time       | raw.TimeUnix       | -                          | -                                   |
