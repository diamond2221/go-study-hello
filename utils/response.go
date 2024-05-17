package utils

import (
	"encoding/json"
	"hello/types"
	"net/http"
	"reflect"
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

func TransformJson(data []byte, reqData *map[string]interface{}) error {
	err := json.Unmarshal(data, &reqData)
	return err
}

/**
 * 将interface{}转换为[]interface{}
 */
func TransformInterfaceToArray(data interface{}, slice *[]interface{}) {
	dataValue := reflect.ValueOf(data)
	if data != nil {
		for i := 0; i < dataValue.Len(); i++ {
			*slice = append(*slice, dataValue.Index(i).Interface())
		}
	}
}
