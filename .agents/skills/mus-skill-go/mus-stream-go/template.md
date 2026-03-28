# Template

Template of the MUS stream serializer for the user provided type. Use predefined 
or another generated MUS stream serializers to process the fields (if the type 
is a struct) or the value itself (if it's a defined type):

```go
var FooMUS = fooMUS{}

type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, w mus.Writer) (n int, err error) {
  ...
}

func (s fooMUS) Unmarshal(r mus.Reader) (v Foo, n int, err error) {
  ...
}

func (s fooMUS) Size(v Foo) (size int) { 
  ...
}

func (s fooMUS) Skip(r mus.Reader) (n int, err error) {
  ...
}
```

`fooMUS` implements `mus.Serializer[Foo]` interface.
