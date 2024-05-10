package api

import (
	"encoding/json"
	"fmt"
	"hello/utils"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 禁止 GET 请求
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		user := strings.Split(r.Form.Get("name"), "")

		w.Write([]byte(fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta http-equiv="X-UA-Compatible" content="IE=edge">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Document</title>
				</head>
				<body>
					<div>%s，您好，当前%s</div>
				</body>
			</html>
		`, strings.Join(user, "_"), "禁止 GET 请求")))
		return
	}
	Origin := r.Header.Get("Origin")
	// 判断Origin是否为空
	if Origin != "" && Origin != "null" {
		// 设置 CORS 头
		w.Header().Set("Access-Control-Allow-Origin", Origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			// 预检请求，直接返回成功
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// 获取body请求参数
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	var reqData map[string]interface{}
	json.Unmarshal(body, &reqData)

	if reqData["error"] != nil || err != nil {
		utils.ErrorResponse(w, 400, "请求参数解析失败", reqData["error"])
	} else {
		var newName string
		if reqData["name"] != nil {
			// 6. 获取请求参数
			// {
			// 	"name": "马超",
			// 	"type": "add"
			// }
			newName = reqData["name"].(string)
		}

		idListObj := reflect.ValueOf(reqData["idList"])
		var idList []interface{}
		if reqData["idList"] != nil {
			for i := 0; i < idListObj.Len(); i++ {
				item := idListObj.Index(i).Interface()
				idList = append(idList, item)
				if reflect.TypeOf(item).Kind() == reflect.Map {
					fmt.Println("map", item.(map[string]interface{})["path"])
				}
			}
		}

		fmt.Println(idList)

		// 将所有参数返回
		data := map[string]interface{}{
			"body": reqData,
		}

		if newName != "" {
			data["name"] = newName
		}
		utils.SuccessJsonResponse(w, data)
	}
}
