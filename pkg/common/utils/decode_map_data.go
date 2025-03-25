package common

import "encoding/json"

func DecodeMapData[T interface{}](encoded interface{}) (T, error) {
	var data T

	result, _ := json.Marshal(encoded)
	err := json.Unmarshal(result, &data)

	return data, err
}
