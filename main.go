package main

import (
	"encoding/json"
	"fmt"
	"hello/api"
	"net/http"
)

func GetPostData(r *http.Request) map[string]interface{} {
	// 1. 获取请求报文的内容长度
	len := r.ContentLength
	// 2. 新建一个字节切片，长度与请求报文的内容长度相同
	body := make([]byte, len)
	// 3. 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	// 4. 将字节切片转换为字符串
	bodyStr := string(body)
	// 5. 将字符串转换为 map
	var reqData map[string]interface{}
	err := json.Unmarshal([]byte(bodyStr), &reqData)
	if err != nil {
		fmt.Println("请求参数解析失败", err)
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	return reqData
}

type User struct {
	id   int
	name string
}

func (u User) Test() {
	fmt.Println(u)
}

func main() {
	u := User{1, "Tom"}
	u1 := u
	u.id = 2
	mValue := u1.Test
	u.Test()
	mValue()
	// mCon := (*User).Test
	// mCon(&u)

	const port = "8080"
	http.HandleFunc("/api/home", api.Home)
	http.HandleFunc("/api/users", api.UserHandler)
	http.HandleFunc("/api/func", api.FuncHandler)
	fmt.Printf("Server is running at: %s%s\n", "http://127.0.0.1:", port)
	http.ListenAndServe(":"+port, nil)
}
