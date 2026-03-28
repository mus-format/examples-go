package main

type FooV2 struct {
	str string
}

func (f FooV2) MarshalMUS(bs []byte) (n int) {
	return FooV2TypedMUS.Marshal(f, bs)
}

func (f FooV2) SizeMUS() (size int) {
	return FooV2TypedMUS.Size(f)
}
