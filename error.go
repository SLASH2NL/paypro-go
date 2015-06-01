package payprogo

type ApiError struct {
	Message string `json:"return"`
	Errors  string `json:"errors"`
}
