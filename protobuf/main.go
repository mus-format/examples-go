package main

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	assert "github.com/ymz-ncnk/assert/panic"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	assert.On = true
}

func main() {
	var (
		dataV1 = DataV1{
			Str:     gofakeit.UUID(),
			Bool:    gofakeit.Bool(),
			Int32:   gofakeit.Int32(),
			Float64: gofakeit.Float64(),
			Slice:   []int32{gofakeit.Int32(), gofakeit.Int32()},
			Time:    timestamppb.New(gofakeit.Date()),
		}
		dataV2 = DataV2{
			Str:     gofakeit.UUID(),
			Int32:   gofakeit.Int32(),
			Float64: gofakeit.Float64(),
			Time:    timestamppb.New(gofakeit.Date()),
		}
	)
	// Marshal using protobuf and unmarshal using mus-go implementation (at
	// the end the unmarshalled data is compared with the original).
	MarshalProtobuf_UnmarshalMusGo(&dataV1)
	// Marshal using mus-go - unmarshal using protobuf.
	MarshalMusGo_UnmarshalProtobuf(&dataV1)
	// Marshal first version and unmarshal second, both using mus-go.
	MarshalDataV1_UnmarshalDataV2(&dataV1)
	// Marshal second version and unmarshal first one again using mus-go.
	MarshalDataV2_UnmarshalDataV1(&dataV2)

	// As you can see, everything works as expected.
}

func MarshalProtobuf_UnmarshalMusGo(data *DataV1) {
	bs, err := proto.Marshal(data)
	assert.EqualError(err, nil)

	adata, _, err := DataV1Protobuf.Unmarshal(bs)
	assert.EqualError(err, nil)

	assert.Equal(data.String(), adata.String())
}

func MarshalMusGo_UnmarshalProtobuf(data *DataV1) {
	bs := make([]byte, DataV1Protobuf.Size(data))
	DataV1Protobuf.Marshal(data, bs)

	adata := DataV1{}
	err := proto.Unmarshal(bs, &adata)
	assert.EqualError(err, nil)

	assert.Equal(data.String(), adata.String())
}

func MarshalDataV1_UnmarshalDataV2(dataV1 *DataV1) {
	bs := make([]byte, DataV1Protobuf.Size(dataV1))
	DataV1Protobuf.Marshal(dataV1, bs)

	dataV2, _, err := DataV2Protobuf.Unmarshal(bs)
	assert.EqualError(err, nil)

	if err := same(dataV1, dataV2); err != nil {
		panic(err)
	}
}

func MarshalDataV2_UnmarshalDataV1(dataV2 *DataV2) {
	bs := make([]byte, DataV2Protobuf.Size(dataV2))
	DataV2Protobuf.Marshal(dataV2, bs)

	dataV1, _, err := DataV1Protobuf.Unmarshal(bs)
	assert.EqualError(err, nil)

	if err := same(dataV1, dataV2); err != nil {
		panic(err)
	}
}

func same(dataV1 *DataV1, dataV2 *DataV2) (err error) {
	if dataV1.Str != dataV2.Str {
		return errors.New("Str")
	}
	if dataV1.Int32 != dataV2.Int32 {
		return errors.New("Int32")
	}
	if dataV1.Float64 != dataV2.Float64 {
		return errors.New("Float64")
	}
	if dataV1.Time != nil && dataV2.Time != nil {
		if dataV1.Time.Seconds != dataV2.Time.Seconds {
			return errors.New("Seconds")
		}
		if dataV1.Time.Nanos != dataV2.Time.Nanos {
			return errors.New("Nanos")
		}
	}
	if (dataV1.Time != nil && dataV2.Time == nil) ||
		(dataV1.Time == nil && dataV2.Time != nil) {
		return errors.New("one time is nil, another is not")
	}
	return
}
