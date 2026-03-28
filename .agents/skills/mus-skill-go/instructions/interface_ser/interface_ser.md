# Interface Serializer

To generate a serializer for an interface type:
1. If there is no `// mus:impls` hint, skip this type.
2. For each implementation type listed in the `// mus:impls` hint:
   - Generate a common serializer (e.g., `FooMUS`).
   - Generate a DTM constants (e.g., `FooDTM`) if needed, see 
     [DTM Constants](./dtm.md).
   - Use DTM constants to create a [Typed Serializers](./typed_ser.md) (e.g., 
    `FooTypedMUS`).
3. Generate the interface serializer that uses the generated typed serializers
   and handles `// mus:marshaller = true` hint.

Example 1, without `// mus:marshaller = true` hint:

```go
// User-provided source code:

const (
  FooDTM com.DTM = 1
  BarDTM com.DTM = 2
)

type Foo struct {...}
func (f Foo) Method() {...}

type Bar struct {...}
func (b Bar) Method() {...}

// mus:impls = Foo, Bar
type Interface interface {
  Method()
}

// -----------------------------------------------------------------------------
// Generated code:

var InterfaceMUS = interfaceMUS{}

type interfaceMUS struct{}

func (s interfaceMUS) Marshal(v Interface, ... ) (n int) {
  switch val := v.(type) {
  case Foo:
    return FooTypedMUS.Marshal(val, ... )
  case Bar:
    return BarTypedMUS.Marshal(val, ... )
  default:
		panic(fmt.Sprintf(com.ErrorPrefix+"unexpected %v type", t))
  }
}

func (s interfaceMUS) Unmarshal( ... ) (v Interface, n int, err error) {
  dtm, n, err := typed.DTMSer.Unmarshal( ... )
  if err != nil {
    return
  }
  switch dtm {
  case FooDTM:
    var v1 Foo
    v1, n1, err1 := FooTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  case BarDTM:
    var v1 Bar
    v1, n1, err1 := BarTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  default:
		err = com.NewUnexpectedDTMError(dtm)
    return
  }
}

func (s interfaceMUS) Size( ... ) (size int) {
	switch t := v.(type) {
	case Foo:
		return FooTypedMUS.Size(t)
	case Bar:
		return BarTypedMUS.Size(t)
	default:
		panic(fmt.Sprintf(com.ErrorPrefix+"unexpected %v type", t))
	}
}

func (s interfaceMUS) Skip( ... ) (n int, err error) {
	dtm, n, err := typed.DTMSer.Unmarshal( ... )
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case FooDTM:
		n1, err = FooTypedMUS.SkipData( ... )
	case BarDTM:
		n1, err = BarTypedMUS.SkipData( ... )
	default:
		err = com.NewUnexpectedDTMError(dtm)
		return
	}
	n += n1
	return
} 
```

Example 2, with `// mus:marshaller = true` hint:

```go
// User-provided source code:

const (
  FooDTM com.DTM = 1
  BarDTM com.DTM = 2
)

type Foo struct {...}
func (f Foo) Method() {...}

type Bar struct {...}
func (b Bar) Method() {...}

// mus:impls = Foo, Bar
// mus:marshaller = true
type Interface interface {
  Method()
}

// -----------------------------------------------------------------------------
// Generated code:

var InterfaceMUS = interfaceMUS{}

type interfaceMUS struct{}

func (s interfaceMUS) Marshal(v Interface, ... ) (n int) {
	if m, ok := v.(mus.MarshallerTyped); ok {
		return m.MarshalTypedMUS(bs)
	}
	panic(fmt.Sprintf("%v doesn't implement the mus.MarshallerTyped interface", reflect.TypeOf(v)))
}

func (s interfaceMUS) Unmarshal( ... ) (v Interface, n int, err error) {
  dtm, n, err := typed.DTMSer.Unmarshal( ... )
  if err != nil {
    return
  }
  switch dtm {
  case FooDTM:
    var v1 Foo
    v1, n1, err1 := FooTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  case BarDTM:
    var v1 Bar
    v1, n1, err1 := BarTypedMUS.UnmarshalData( ... )
    return v1, n + n1, err1
  default:
		err = com.NewUnexpectedDTMError(dtm)
    return
  }
}

func (s interfaceMUS) Size( ... ) (size int) {
	if m, ok := v.(mus.MarshallerTyped); ok {
		return m.SizeTypedMUS()
	}
	panic(fmt.Sprintf("%v doesn't implement the mus.MarshallerTyped interface", reflect.TypeOf(v)))
}

func (s interfaceMUS) Skip( ... ) (n int, err error) {
	dtm, n, err := typed.DTMSer.Unmarshal( ... )
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case FooDTM:
		n1, err = FooTypedMUS.SkipData( ... )
	case BarDTM:
		n1, err = BarTypedMUS.SkipData( ... )
	default:
		err = com.NewUnexpectedDTMError(dtm)
		return
	}
	n += n1
	return
} 
```

In this case the user is responsible for implementing `mus.MarshallerTyped` 
interface by each implementation type.
