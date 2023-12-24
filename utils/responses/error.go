package responses

type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

func ErrorResponse(message string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}

