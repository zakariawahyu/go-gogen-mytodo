package payload

import (
	"zakariawahyu.com/go-gogen-mytodo/shared/model/apperror"
)

type SuccessResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	TraceID string `json:"traceId"`
}

type ErrorResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage any    `json:"errorMessage"`
	TraceID      string `json:"traceId"`
}

func NewSuccessResponse(data any, traceID string) any {
	var res SuccessResponse
	res.Success = true
	res.Data = data
	res.TraceID = traceID
	return res
}

func NewErrorResponse(err error, traceID string) any {
	var res ErrorResponse
	res.Success = false
	res.TraceID = traceID

	et, ok := err.(apperror.ErrorType)
	if !ok {
		res.ErrorCode = "UNDEFINED"
		res.ErrorMessage = err.Error()
		return res
	}

	res.ErrorCode = et.Code()
	res.ErrorMessage = et.Error()
	return res
}
