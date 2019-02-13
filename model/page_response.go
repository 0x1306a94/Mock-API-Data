package model

type PageResponse struct {
	PageNo   int64       `json:"pageNo"`
	PageSize int64       `json:"pageSize"`
	Total    int64       `json:"total"`
	HasMore  bool        `json:"hasMore"`
	List     interface{} `json:"list"`
}
