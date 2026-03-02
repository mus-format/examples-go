# Validation Examples

This directory demonstrates how to perform data validation during unmarshalling 
with the `mus-go` library.

- [string](string/main.go): Shows how to use `ord.NewValidStringSer` with a 
  length validator.
- [slise](slise/main.go): Shows how to use `ord.NewValidSliceSer` with length 
  and element validators.
- [map](map/main.go): Shows how to use `ord.NewValidMapSer` with length, key, 
  and value validators.

Validation is integrated into the unmarshalling process, allowing for early 
failure if the data does not meet required criteria.