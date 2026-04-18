# Marshal Function Example

This example demonstrates how to use the `mus.Marshal` function.

Instead of manual buffer management:
```go
bs = make([]byte, v.SizeMUS())
v.MarshalMUS(bs)
```

You can use:
```go
bs := mus.Marshal(v)
```
