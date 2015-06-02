package payprogo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ApiError struct {
	Message string `json:"return"`
	Errors  bool   `json:"errors"`
}

type StringBoolean bool

func (e *ApiError) UnmarshalJSON(b []byte) error {
	var intermediate struct {
		*ApiError
		Errors string
	}

	err := json.Unmarshal(b, &intermediate)
	if err != nil {
		return err
	}

	val, err := strconv.ParseBool(intermediate.Errors)
	if err != nil {
		return fmt.Errorf("Expected 'false' or 'true' got:%s", string(b))
	}

	e.Errors = val
	e.Message = intermediate.Message

	return nil
}
