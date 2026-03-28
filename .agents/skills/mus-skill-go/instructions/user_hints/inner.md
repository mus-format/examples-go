# Inner Hints

## Number Encoding

For all number types the user can specify the encoding.

```go
// mus:numEnc = varint | raw
int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

// mus:numEnc = positive
int, int8, int16, int32, int64
```

Where:

- mus:numEnc = varint -> varint.Int
- mus:numEnc = positive -> varint.PositiveInt
- mus:numEnc = raw -> raw.Int

Example:

```go
// mus:numEnc = raw
type Foo int

type fooSer struct{}

func (s fooSer) Marshal( ... ) (n int) {
	return raw.Int.Marshal( ... )
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v1, n, err = raw.Int.Unmarshal( ... )
	if err != nil {
		return
	}
	v = Foo(v1)
	return
}

func (s fooSer) Size(v Foo) (size int) {
	return raw.Int.Size(int(v))
}

func (s fooSer) Skip( ... ) (n int, err error) {
	return raw.Int.Skip( ... )
}
```

## Time

```go
// mus:time_unit = ...
time.Time
```

Where:

- mus:time_unit = sec -> TimeUnix
- mus:time_unit = milli -> TimeUnixMilli
- mus:time_unit = micro -> TimeUnixMicro
- mus:time_unit = nano -> TimeUnixNano
- mus:time_unit = sec_utc -> TimeUnixUTC
- mus:time_unit = milli_utc -> TimeUnixMilliUTC
- mus:time_unit = micro_utc -> TimeUnixMicroUTC
- mus:time_unit = nano_utc -> TimeUnixNanoUTC

Example:

```go
// mus:time_unit = milli
type Foo time.Time

type fooSer struct{}

func (s fooSer) Marshal(v Foo, ... ) (n int) {
	return raw.TimeUnixMilli.Marshal(time.Time(v), ... )
}
...
```

## Field Validator

```go
// mus:vl = ...
T
```

Where:

- vl: `com.Validator[T]` interface implementation

Example:

```go
type Foo struct {
  // mus:vl = ValidateNum
	num int
	// mus:vl = ValidateStr
	str string
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	var v1 int
	v.num, n, err = varint.Int.Unmarshal( ... )
	if err != nil {
		return
	}
	if err = ValidateNum(v.num); err != nil {
		return
	}
	var n1 int
	v.str, n1, err = ord.String.Unmarshal( ... )
	n += n1
	if err != nil {
		return
	}
	err = ValidateStr(v.str)
	return
}
```

## Length Encoding

```go
// mus:lenEnc = ...
string, array, slice or map
```

Where:

- mus:lenEnc = varint -> varint.Int
- mus:lenEnc = positive -> varint.PositiveInt
- mus:lenEnc = raw -> raw.Int

## String

String type hints require creating and using an anonymous serializer for that 
type. 

```go
// mus:lenEnc = ...
// mus:lenVl = ...
string
```

Where:

- mus:lenVl: `com.Validator[int]` interface implementation

Example:

```go
// mus:lenEnc = raw
type Foo string

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewStringSer(strops.WithLenSer(raw.Int))
```

Example valid:

```go
// mus:lenEnc = raw
// mus:lenVl: ValidateLength
type Foo string

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var str`generated hash value` = ord.NewValidStringSer(
	stropts.WithLenSer(raw.Int), 
  stropts.WithLenValidator(ValidateLength),
)
```

## Array

```go
// mus:lenEnc = ...
// mus:elemVl = ...
[N]T
```

Where:

- mus:elemVl: `com.Validator[T]` interface implementation

Example:

```go
// mus:lenEnc = raw
type ArrayInt [3]int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var array`generated hash value` = unsafe.NewArraySer(varint.Int, 
  arropts.WithLenSer[int](raw.Int), 
)
```

Example valid:

```go
// mus:lenEnc = raw
// mus:elemVl: ValidateElement
type ArrayInt [3]int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var array`generated hash value` = unsafe.NewValidArraySer(varint.Int, 
  arropts.WithLenSer[int](raw.Int), 
  arropts.WithElemValidator[int](ValidateElement),
)
```

## Slice

```go
// mus:lenEnc = ...
// mus:lenVl = ...
// mus:elemVl = ...
[]T
```

Where:

- mus:lenVl: `com.Validator[int]` interface implementation
- mus:elemVl: `com.Validator[T]` interface implementation

Example:

```go
// mus:lenEnc = raw
type SliceInt []int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var slice`generated hash value` = ord.NewSliceSer(varint.Int, 
  slopts.WithLenSer[int](raw.Int), 
)
```

Example valid:

```go
// mus:lenEnc = raw
// mus:lenVl: ValidateLength
// mus:elemVl: ValidateElement
type SliceInt []int

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var slice`generated hash value` = ord.NewValidSliceSer(varint.Int, 
  slopts.WithLenSer[int](raw.Int), 
	slopts.WithLenValidator[int](ValidateLength), 
  slopts.WithElemValidator[int](ValidateElement),
)
```

## Byte Slice

```go
// mus:lenEnc = ...
// mus:lenVl = ...
[]byte
```
Where:

- lenVl: `com.Validator[int]` interface implementation

Example:

```go
// mus:lenEnc = raw
type ByteSlice []byte 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var bs`generated hash value` = ord.NewByteSliceSer(bslops.WithLenSer(raw.Int))
```

Example valid:

```go
// mus:lenEnc = raw
// mus:lenVl: ValidateLength
type ByteSlice []byte 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var bs`generated hash value` = ord.NewValidByteSliceSer(
	bslopts.WithLenSer(raw.Int), 
  bslopts.WithLenValidator(ValidateLength),
)
```

## Map

```go
// mus:lenEnc = ...
// mus:lenVl = ...
// mus:keyVl = ...
// mus:elemVl = ...
map[K]V
```

Where:

- mus:lenVl: `com.Validator[int]` interface implementation
- mus:keyVl: `com.Validator[K]` interface implementation
- mus:elemVl: `com.Validator[V]` interface implementation

Example:

```go
// mus:lenEnc = raw
type Map map[string]int 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var map`generated hash value` = ord.NewMapSer(ord.String, varint.Int,
	mapopts.WithLenSer[string, int](raw.Int),
)
```

Example valid:

```go
// mus:lenEnc = raw
// mus:lenVl: ValidateLength
// mus:keyVl: ValidateKey
// mus:elemVl: ValidateValue
type Map map[string]int 

// Marshal, Unmarshal, Size, Skip methods use this serializer.
var map`generated hash value` = ord.NewValidMapSer(ord.String, varint.Int,
	mapopts.WithLenSer[string, int](varint.Int),
	mapopts.WithLenValidator[string, int](ValidateLength),
	mapopts.WithKeyValidator[string, int](ValidateKey),
	mapopts.WithValueValidator[string, int](ValidateValue),
)
```

## Ignore Field

With the `mus:ignore` hint, the user can skip a struct field from the serialization 
process.

```go
// mus:ignore = true
```

Example:

```go
type Foo struct {
  num int
  // mus:ignore = true
  secret string
}

func (s fooSer) Marshal(v Foo, ... ) (n int) {
	return varint.Int.Marshal(v.num, ...)
}

func (s fooSer) Unmarshal( ... ) (v Foo, n int, err error) {
	v.num, n, err = varint.Int.Unmarshal( ... )
	return
}

func (s fooSer) Size(v Foo) (size int) {
	return varint.Int.Size(v.num)
}

func (s fooSer) Skip( ... ) (n int, err error) {
	return varint.Int.Skip( ... )
}
```