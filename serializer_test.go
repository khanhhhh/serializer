package serializer

import (
	"fmt"
	"testing"
)

func TestSerializer(t *testing.T) {
	bytes, err := Marshal(map[string]interface{}{
		"name": "Khanh",
		"age":  23,
		"file": []byte{1, 2, 3},
		"data": map[string]interface{}{
			"key": -1234,
		},
	})
	if err != nil {
		panic(err)
	}
	object, err := Unmarshal(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", object)
}
