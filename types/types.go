package types

type Integer interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | int
}

type SignedInteger interface {
	int8 | int16 | int32 | int64 | int
}

type UnsignedInteger interface {
	uint8 | uint16 | uint32 | uint64
}
