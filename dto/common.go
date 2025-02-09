package dto

import "time"

type BaseEntity struct {
	UserId  string `json:"userId"`
	TraceId string `json:"traceId"`
	AppId   string `json:"appId"`
}

type ModelQueryRequest struct {
	Base      *BaseEntity `json:"base"`
	Question  string      `json:"question"`
	ModelType string      `json:"modelType"`
}

type ModelQueryResponse struct {
	Response string `json:"response"`
}

type CommonResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:      0,
		Message:   "success",
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}

func NewFailResponse(code int, message string) *CommonResponse {
	return &CommonResponse{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().UnixMilli(),
		Data:      nil,
	}
}
