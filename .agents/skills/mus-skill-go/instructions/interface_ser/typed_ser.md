# Typed Serializer

Consist of DTM and common serializer.

## User Expectations

The user may define a DTM (Data Type Metadata) constant for each type that 
requires a typed serializer.

By convention, for a type `T`, the DTM constant should be named `TDTM`. 
The generator uses this constant to create the typed serializer.

## Generation Steps

To generate a typed serializer:
1. Generate a common serializer for a type if one doesn't already exist.
2. Generate a DTM constant if one doesn't already exist.
3. Use `typed.NewTypedSer` to create a typed serializer.

Example:

```go
// The user defined DTM constant.
const FooDTM com.DTM = 1 

type Foo struct {...}

var FooMUS = fooSer{}

var FooTypedMUS = typed.NewTypedSer(FooDTM, FooMUS)
```
