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
