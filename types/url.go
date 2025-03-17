package types

type Url struct {
	Code        string `json:"code"`
	OriginalUrl string `json:"original_url"`
	VisitCount  int64  `json:"visit_count"`
}

type CreateUrlPayload struct {
	OriginalUrl string `json:"original_url" validate:"required,url"`
}