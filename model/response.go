package model

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ResponseSuccess struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
