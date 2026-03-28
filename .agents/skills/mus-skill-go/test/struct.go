package test

import (
	"time"
)

type SimpleStruct struct {
	Num int
	Str string
}

type IgnoreStruct struct {
	Num int
	// mus:ignore = true
	Str string
}

type TimeStruct time.Time

type EmbeddingStruct struct {
	Num int
	InnerStruct
}

type InnerStruct struct {
	Str string
}

// mus:vl = ValidateStruct
type ValidStruct struct {
	// mus:vl = ValidateNum
	Num int
	// mus:vl = ValidateStr
	Str string
}

type ParametricStruct[T any] struct {
	Field T
}
