package responses

type TSuccessResponsePage struct {
	Meta    TResponseMetaPage `json:"meta"`
	Results interface{}       `json:"results"`
}

func SuccessResponsePage(message string, page int, limit int, total int64, data interface{}) interface{} {
	return TSuccessResponsePage{
		Meta: TResponseMetaPage{
			Success: true,
			Message: message,
			Page:    page,
			Limit:   limit,
			Total:   total,
		},
		Results: data,
	}
}

