package helpers

type Response struct {
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Claims  interface{} `json:"claims,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
