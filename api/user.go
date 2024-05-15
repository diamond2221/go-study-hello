package api

import (
	"encoding/json"
	"hello/utils"
	"io"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	AddTime string `json:"add_time"`
	UpTime  string `json:"up_time"`
	UpdTime string `json:"upd_time"`
	Todos   []Todo `json:"todos" gorm:"foreignKey:AddUser;references:Id"`
}

// TableName 方法指定表名称
func (User) TableName() string {
	return "tbl_user" // 自定义表名称为 tbl_user
}

type Todo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	AddTime string `json:"add_time"`
	AddUser string `json:"add_user"`
	Status  int    `json:"status"`
	User    User   `json:"user" gorm:"foreignKey:AddUser;references:Id"`
}

func (Todo) TableName() string {
	return "tbl_todo"
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		CreateUser(w, r)
	} else if r.Method == "GET" {
		GetUsers(w, r)
	}
}
func Concent() *gorm.DB {
	var dsn = "root:981220Zy@tcp(127.0.0.1:3306)/db_todo?charset=utf8mb4&parseTime=True&loc=Local"
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	page, _ := strconv.Atoi(r.Form.Get("page"))
	limit, _ := strconv.Atoi(r.Form.Get("limit"))

	var users []struct {
		User
		Count int
	}

	db := Concent()

	db.Preload("Todos").Limit(limit).Offset((page-1)*limit).Where("name like ?", "%"+r.Form.Get("name")+"%").Find(&users)

	for i := 0; i < len(users); i++ {
		users[i].Count = len(users[i].Todos)
	}
	utils.SuccessJsonResponse(w, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var data map[string]string

	json.Unmarshal(body, &data)

	db := Concent()

	now := time.Now().Format("2006/01/02 15:04:05")
	user := User{
		Name:    data["name"],
		Phone:   data["phone"],
		AddTime: now,
		UpTime:  now,
		UpdTime: now,
	}

	var hasUser User
	db.Where("name = ?", user.Name).First(&hasUser)
	if hasUser.Id == 0 {
		db.Create(&user)
		utils.SuccessJsonResponse(w, user)
	} else {
		utils.ErrorResponse(w, 400, "用户: "+hasUser.Name+",已存在", nil)
	}

}
