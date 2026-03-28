# Predefined Constructors

Constructors are used to create serializers with additional options.

## ord

Contains constructors for the following types:

- `array`: 
  unsafe.NewArraySer[T any](elemSer mus.Serializer[T], opts ...arropts.SetOption[T]) arraySer[T]
  unsafe.NewValidArraySer[T any](elemSer mus.Serializer[T], opts ...arropts.SetOption[T]) validArraySer[T]
- `[]byte`: 
  ord.NewByteSliceSer(opts ...bslops.SetOption) byteSliceSer 
  ord.NewValidByteSliceSer(opts ...bslops.SetOption) validByteSliceSer 
  unsafe.NewByteSliceSer(opts ...bslops.SetOption) byteSliceSer 
  unsafe.NewValidByteSliceSer(opts ...bslops.SetOption) validByteSliceSer 
- `map`: 
  ord.NewMapSer[T comparable, V any](keySer mus.Serializer[T], valueSer mus.Serializer[V], opts ...mapops.SetOption[T, V]) mapSer[T, V]
  ord.NewValidMapSer[T comparable, V any](keySer mus.Serializer[T], valueSer mus.Serializer[V], opts ...mapops.SetOption[T, V]) validMapSer[T, V] 
- `pointer`: 
  ord.NewPtrSer[T any](baseSer mus.Serializer[T]) ptrSer[T] 
- `slice`: 
  ord.NewSliceSer[T any](elemSer mus.Serializer[T], opts ...slops.SetOption[T]) sliceSer[T]
	NewValidSliceSer[T any](elemSer mus.Serializer[T], opts ...slops.SetOption[T]) validSliceSer[T]
- `string`: 
  ord.NewStringSer(opts ...strops.SetOption) stringSer
  ord.NewValidStringSer(opts ...strops.SetOption) validStringSer


## unsafe

Contains constructors for the following types:

- `string`:
  NewStringSer(opts ...stropts.SetOption) stringSer
  NewValidStringSer(opts ...stropts.SetOption) validStringSer
