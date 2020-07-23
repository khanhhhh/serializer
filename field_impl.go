package serializer

import (
	"encoding/binary"
	"errors"
)

func toField(attribute string, data interface{}) (field field, err error) {
	// attribute
	field.attribute = []byte(attribute)
	field.attributeLength = uint64(len(field.attribute))
	// data
	switch data.(type) {
	case int:
		field.dataType = integerType
		field.data = make([]byte, 8)
		binary.LittleEndian.PutUint64(field.data, uint64(data.(int)))
	case string:
		field.dataType = stringType
		field.data = []byte(data.(string))
	case []byte:
		field.dataType = byteArrayType
		field.data = []byte(data.([]byte))
	case map[string]interface{}:
		field.dataType = objectType
		field.data, err = Marshal(data.(map[string]interface{}))
		if err != nil {
			return field, err
		}
	default:
		err := errors.New("Field is not: int, string, []byte")
		return field, err
	}
	field.dataLength = uint64(len(field.data))
	return field, nil
}

func fromField(field field) (attribute string, data interface{}, err error) {
	attribute = string(field.attribute)
	switch field.dataType {
	case integerType:
		data = int(binary.LittleEndian.Uint64(field.data))
	case stringType:
		data = string(field.data)
	case byteArrayType:
		data = make([]byte, len(field.data))
		copy(data.([]byte), field.data)
	case objectType:
		data, err = Unmarshal(field.data)
		if err != nil {
			return attribute, data, err
		}
	}
	return attribute, data, nil
}
func (field field) toBytes() (bytes []byte) {
	bytes = make([]byte, 17)
	binary.LittleEndian.PutUint64(
		bytes[0:8],
		field.attributeLength,
	)
	bytes[8] = field.dataType
	binary.LittleEndian.PutUint64(
		bytes[9:17],
		field.dataLength,
	)
	bytes = append(bytes, field.attribute...)
	bytes = append(bytes, field.data...)
	return bytes
}

func fromBytes(bytes []byte) (fields []field, err error) {
	fields = []field{}
	if len(bytes) == 0 {
		return []field{}, nil
	}
	for len(bytes) > 0 {
		if len(bytes) < 17 {
			return fields, errors.New("Unmarshal Error")
		}
		newField := field{
			attributeLength: binary.LittleEndian.Uint64(bytes[0:8]),
			dataType:        bytes[8],
			dataLength:      binary.LittleEndian.Uint64(bytes[9:17]),
			attribute:       nil,
			data:            nil,
		}
		if len(bytes) < int(17+newField.attributeLength+newField.dataLength) {
			return fields, errors.New("Unmarshal Error")
		}
		newField.attribute = make([]byte, newField.attributeLength)
		newField.data = make([]byte, newField.dataLength)
		copy(
			newField.attribute,
			bytes[17:17+newField.attributeLength],
		)
		copy(
			newField.data,
			bytes[17+newField.attributeLength:17+newField.attributeLength+newField.dataLength],
		)
		fields = append(fields, newField)
		//////
		bytes = bytes[17+newField.attributeLength+newField.dataLength:]
	}
	return fields, nil
}
