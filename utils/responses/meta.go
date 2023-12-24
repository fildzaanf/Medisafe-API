package responses

type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TResponseMetaPage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Total   int64  `json:"total"`
}
