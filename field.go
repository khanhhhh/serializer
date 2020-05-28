package serializer

// DataTypeConstants :
const (
	integerType   = 0
	stringType    = 1
	byteArrayType = 2
	objectType    = 3
)

// Type :
type dataType = uint8

// field :
type field struct {
	attributeLength uint64   // 0 -> 8
	dataType        dataType // 8 -> 9
	dataLength      uint64   // 9 -> 17
	attribute       []byte
	data            []byte
}
