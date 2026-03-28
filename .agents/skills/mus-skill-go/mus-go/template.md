# Template

Template of the MUS serializer for the user provided type. Use predefined or 
another generated MUS serializers to process the fields (if the type is 
a struct) or the value itself (if it's a defined type):

```go
var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  ...
}

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
  ...
}

func (s fooMUS) Size(v Foo) (size int) { 
  ...
}

func (s fooMUS) Skip(bs []byte) (n int, err error) {
  ...
}
```

`fooMUS` implements `mus.Serializer[Foo]` interface.
