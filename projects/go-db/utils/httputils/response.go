package httputils

type ResponseTemplate struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func BuildResponse(data interface{}) interface{} {
	return ResponseTemplate{true, data}
}
