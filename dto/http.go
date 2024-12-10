package dto

type Response struct {
	Error *string `json:"error,omitempty"`
	Data  any     `json:"data,omitempty"`
}

type Error string

func Point[T any](data T) *T {
	return &data
}
