# Unsafe Serialization

This example demonstrates how to use the `unsafe` package.

For strings, `unsafe.String.Unmarshal` creates a string that points directly 
to the underlying byte slice. This avoids memory allocation and copying but 
means the string's content will change if the byte slice is modified.