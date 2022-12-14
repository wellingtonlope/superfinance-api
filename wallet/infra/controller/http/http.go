package http

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	if err == nil {
		return ErrorResponse{}
	}
	return ErrorResponse{err.Error()}
}
