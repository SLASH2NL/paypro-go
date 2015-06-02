package payprogo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type apiError struct {
	Message string `json:"return"`
	Errors  bool   `json:"errors"`
}

// Avoid recursion
type _error apiError

func (e *apiError) UnmarshalJSON(b []byte) error {
	var intermediate struct {
		_error
		Errors string `json:"errors"`
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
