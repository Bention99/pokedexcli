package api

import (
	"encoding/json"
)

func unmarshal[T any](b []byte) (T, error) {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		var zero T
		return zero, err
	}
	return v, nil
}