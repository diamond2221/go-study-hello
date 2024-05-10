package utils

import (
	"encoding/json"
	"hello/types"
	"net/http"
)

func SuccessJsonResponse(w http.ResponseWriter, data interface{}) {
	var response = types.Res{
		Data:    data,
		Message: "success",
		Code:    200,
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}

func ErrorResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	var response = types.Res{
		Data:    data,
		Message: message,
		Code:    code,
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}
