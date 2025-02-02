package column

import (
	"github.com/pqx0703/clickhouse-go/lib/binary"
)

type UInt16 struct{ base }

func (UInt16) Read(decoder *binary.Decoder, isNull bool) (interface{}, error) {
	v, err := decoder.UInt16()
	if err != nil {
		return uint16(0), err
	}
	return v, nil
}

func (u *UInt16) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case uint16:
		return encoder.UInt16(v)
	case int64:
		return encoder.UInt16(uint16(v))
	case uint64:
		return encoder.UInt16(uint16(v))
	case int:
		return encoder.UInt16(uint16(v))

	// this relies on Nullable never sending nil values through
	case *uint16:
		return encoder.UInt16(*v)
	case *int64:
		return encoder.UInt16(uint16(*v))
	case *uint64:
		return encoder.UInt16(uint16(*v))
	case *int:
		return encoder.UInt16(uint16(*v))
	}

	return &ErrUnexpectedType{
		T:      v,
		Column: u,
	}
}
