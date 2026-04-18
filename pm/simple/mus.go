package main

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/pm"
)

// MakeTwoPtrsSer creates a new TwoPtrs serializer that can be directly used.
func MakeTwoPtrsSer() mus.Serializer[TwoPtrs] {
	var (
		ptrMap    = com.NewPtrMap()
		revPtrMap = com.NewReversePtrMap()
		ser       = NewTwoPtrsSer(ptrMap, revPtrMap)
	)
	return pm.Wrap(ptrMap, revPtrMap, ser)
}

func MakeEmbeddedPtrsSer() mus.Serializer[EmbeddedPtrs] {
	var (
		ptrMap    = com.NewPtrMap()
		revPtrMap = com.NewReversePtrMap()
		ser       = NewEmbeddedPtrsSer(ptrMap, revPtrMap)
	)
	return pm.Wrap(ptrMap, revPtrMap, ser)
}

// NewTwoPtrsSer creates a new TwoPtrs serializer.
func NewTwoPtrsSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap) twoPtrsSer {
	strPtrSer := pm.NewPtrSer(ptrMap, revPtrMap, ord.String)
	return twoPtrsSer{strPtrSer}
}

func NewEmbeddedPtrsSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap) mus.Serializer[EmbeddedPtrs] {
	twoSer := NewTwoPtrsSer(ptrMap, revPtrMap)
	strPtrSer := pm.NewPtrSer(ptrMap, revPtrMap, ord.String)
	return embeddedPtrsSer{twoSer, strPtrSer}
}

// twoPtrsSer implements mus.Serializer for TwoPtrs.
type twoPtrsSer struct {
	strPtrSer mus.Serializer[*string]
}

func (s twoPtrsSer) Marshal(p TwoPtrs, bs []byte) (n int) {
	n = s.strPtrSer.Marshal(p.ptr1, bs)
	return n + s.strPtrSer.Marshal(p.ptr2, bs[n:])
}

func (s twoPtrsSer) Unmarshal(bs []byte) (p TwoPtrs, n int, err error) {
	p.ptr1, n, err = s.strPtrSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	p.ptr2, n1, err = s.strPtrSer.Unmarshal(bs[n:])
	n += n1
	return
}

func (s twoPtrsSer) Size(p TwoPtrs) (size int) {
	size = s.strPtrSer.Size(p.ptr1)
	return size + s.strPtrSer.Size(p.ptr2)

}

func (s twoPtrsSer) Skip(bs []byte) (n int, err error) {
	n, err = s.strPtrSer.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.strPtrSer.Skip(bs[n:])
	n += n1
	return
}

// embeddedPtrsSer implements mus.Serializer for EmbeddedPtrs.
type embeddedPtrsSer struct {
	twoSer    twoPtrsSer
	strPtrSer mus.Serializer[*string]
}

func (s embeddedPtrsSer) Marshal(p EmbeddedPtrs, bs []byte) (n int) {
	n = s.twoSer.Marshal(p.TwoPtrs, bs)
	return n + s.strPtrSer.Marshal(p.ptr3, bs[n:])
}

func (s embeddedPtrsSer) Unmarshal(bs []byte) (p EmbeddedPtrs, n int, err error) {
	p.TwoPtrs, n, err = s.twoSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	p.ptr3, n1, err = s.strPtrSer.Unmarshal(bs[n:])
	n += n1
	return
}

func (s embeddedPtrsSer) Size(p EmbeddedPtrs) (size int) {
	size = s.twoSer.Size(p.TwoPtrs)
	return size + s.strPtrSer.Size(p.ptr3)
}

func (s embeddedPtrsSer) Skip(bs []byte) (n int, err error) {
	n, err = s.twoSer.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.strPtrSer.Skip(bs[n:])
	n += n1
	return
}
