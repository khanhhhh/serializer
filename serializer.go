package serializer

// Marshal : Marshal a map of string keys and (int, string, []bytes, map) values
func Marshal(object map[string]interface{}) (bytes []byte, err error) {
	for attribute, data := range object {
		field, err := toField(attribute, data)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, field.toBytes()...)
	}
	return bytes, nil
}

// Unmarshal : Unmarshal
func Unmarshal(bytes []byte) (object map[string]interface{}, err error) {
	fields, err := fromBytes(bytes)
	if err != nil {
		return nil, err
	}
	object = make(map[string]interface{})
	for _, field := range fields {
		attribute, data, err := fromField(field)
		if err != nil {
			return object, err
		}
		object[attribute] = data
	}
	return object, nil
}
