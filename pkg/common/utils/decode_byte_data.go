package common

import "encoding/json"

func DecodeByteData[T interface{}](encoded []byte) (T, error) {
	var data T

	err := json.Unmarshal(encoded, &data)

	return data, err
}
