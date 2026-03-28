package test

// mus:lenEnc = raw
type LenString string

// mus:lenEnc = raw
type LenSlice []int

// mus:lenEnc = raw
type LenArray [3]int

// mus:lenEnc = raw
type LenMap map[string]int

type Slice []int

type Array [3]int

type Map map[string]int

type Ptr *int

// mus:vl = ValidateLen
type ValidString string

// mus:vl = ValidateLen
// mus:elemVl = ValidateNum
type ValidSlice []int

// mus:elemVl = ValidateNum
type ValidArray [3]int

// mus:vl = ValidateLen
// mus:keyVl = ValidateStr
// mus:elemVl = ValidateNum
type ValidMap map[string]int

// mus:vl = ValidateLen
// mus:elemVl = ValidateComplexMapValue
type ComplexMap map[string]map[*[]int][][]float32
