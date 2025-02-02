package column

import (
	"github.com/pqx0703/clickhouse-go/lib/binary"
)

type Float64 struct{ base }

func (Float64) Read(decoder *binary.Decoder, isNull bool) (interface{}, error) {
	v, err := decoder.Float64()
	if err != nil {
		return float64(0), err
	}
	return v, nil
}

func (float *Float64) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case float32:
		return encoder.Float64(float64(v))
	case float64:
		return encoder.Float64(v)

	// this relies on Nullable never sending nil values through
	case *float32:
		return encoder.Float64(float64(*v))
	case *float64:
		return encoder.Float64(*v)
	}

	return &ErrUnexpectedType{
		T:      v,
		Column: float,
	}
}
