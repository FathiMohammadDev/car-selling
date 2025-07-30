package common

import "encoding/json"

func TypeConverter[T any](data any) (*T, error) {
	var result T

	jsonData, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
