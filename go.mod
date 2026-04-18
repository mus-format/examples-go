module github.com/mus-format/examples-go

go 1.24.0

toolchain go1.24.1

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/mus-format/common-go v0.0.0-20260324174526-3d8f1741b5a2
	github.com/mus-format/ext-protobuf-go v0.0.0-20260328200958-82e3c3e92c35
	github.com/mus-format/mus-gen-go v0.5.1
	github.com/mus-format/mus-go v0.10.1
	github.com/ymz-ncnk/assert v0.0.0-20260108210721-155bc9aa4282
	github.com/ymz-ncnk/mok v0.2.2
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/tools v0.41.0 // indirect
)

replace github.com/mus-format/mus-stream-go => ../mus-stream-go
