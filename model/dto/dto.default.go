package dto

type DefaultDTOResponseWrapper struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Body    interface{} `json:"data"`
}

type NoBodyDTOResponseWrapper struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

type PaginationDTOWrapper struct {
	PageSize  int16       `json:"page_size"`
	Page      int16       `json:"page"`
	TotalPage int16       `json:"total_page"`
	Data      interface{} `json:"page_data"`
}
