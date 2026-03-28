# Anonymous Serializers

For a `string` (with a hint), pointer, array, slice or a map type generate an
anonymous serializer.

Anonymous serializer should be created with a constructor from 
[Predefined Constructors](./predefined_constructors.md) and declared as a variable.

Anonymous serializer variable name is build like:

```
// here "+" sign denotes a concatenation

// for a string serializer
string + Base64 (Hash (lenEnc + lenVl))

// for an array serializer
array + Base64 (Hash (elem_type + lenEnc + elemVl))

// for a byte slice serializer
byteSlice + Base64 (Hash (lenEnc + lenVl))

// for a slice serializer
slice + Base64 (Hash (elem_type + lenEnc + lenVl + elemVl))

// for a map serializer
map + Base64 (Hash (key_type + elem_type + lenEnc + lenVl + keyVl + valVl))

// for a pointer serializer
ptr + Base64 (Hash (elem_type))
```

To create valid Go identifiers, the standard Base64 characters MUST be replaced 
as follows:
- `+` -> `Δ`
- `/` -> `Σ`
- `=` -> `Ξ`

Use hash (MD5) in a name to ensure unique names across the package. For
example:

```go
var (
	mapAUEFΣr5wWΣsQKmdS5ns98QΞΞ   = ord.NewMapSer(ord.String, varint.Int64)
	ptry4kBΔSqZ1hMnnFP3Y5bQ4wΞΞ   = ord.NewPtrSer(ord.Bool)
	arraykhqoj8H1iZDaFGbnRRNCSQΞΞ = unsafe.NewArraySer[[10]int](varint.Int)
	sliceoVLU4RNiPduaSOBCraZgCwΞΞ = ord.NewSliceSer(varint.Uint16)
)
```

### Reference Implementation (Python)

To ensure consistency in hash generation, use the following Python script 
[hash_gen.py](../../scripts/hash_gen.py).
