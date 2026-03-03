# Validation Examples

This directory demonstrates how to perform data validation during unmarshalling.

- [string](string/main.go): Use `ord.NewValidStringSer` with a length validator.
- [slice](slice/main.go): Use `ord.NewValidSliceSer` with length and element validators.
- [map](map/main.go): Use `ord.NewValidMapSer` with length, key, and value validators.

Validation is integrated into the unmarshalling process, allowing for early 
failure if the data does not meet required criteria.