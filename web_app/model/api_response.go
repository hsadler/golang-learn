package model

// APIResponse : struct for returning any payload type in an API response
type APIResponse struct {
	Success bool
	Payload interface{}
}
