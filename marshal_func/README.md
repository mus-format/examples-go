# Marshal Function

This example demonstrates how to use the `MarshalMUS` function from the 
[ext-go](https://github.com/mus-format/ext-go) library.

Instead of manual buffer management:
```go
bs = make([]byte, v.SizeMUS())
v.MarshalMUS(bs)
```

You can use:
```go
bs := ext.MarshalMUS(v)
```
