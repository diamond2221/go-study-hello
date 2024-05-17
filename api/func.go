package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClassesId int

func (c ClassesId) getName() string {
	var names = []string{"一班", "二班", "三班", "四班", "五班"}

	if (int(c) > len(names)-1) || (c-1 < 0) {
		return "未知班级"
	}
	name := names[c-1]
	return name
}

type ClassItem struct {
	Name    string    `json:"name"`
	Id      ClassesId `json:"id"`
	Address string    `json:"address"`
}

type Classes []ClassItem

func (c *ClassItem) getClassesName() {
	c.Name = c.Id.getName()
}

type Skill struct {
	SkillName string `json:"skillName"`
	Action    string `json:"action"`
}

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     bool   `json:"sex"`
	SexStr  string `json:"sexStr"`
	Classes `json:"classes"`
	*Skill  `json:"skill"`
}

func (s *student) setSexStr() {
	if s.Sex {
		s.SexStr = "男"
	} else {
		s.SexStr = "女"
	}
}

func FuncHandler(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]*student)
	stus := []student{
		{
			Name: "mayun",
			Age:  18,
			Sex:  true,
			Classes: Classes{
				{Id: 0, Address: "深圳"},
			},
			Skill: &Skill{
				SkillName: "拳击",
				Action:    "打人",
			},
		},
		{
			Name: "wangxiaoli",
			Age:  23,
			Classes: Classes{
				{Id: 2, Address: "北京"},
			},
		},
		{
			Name: "zhangfei",
			Age:  28,
			Sex:  true,
			Classes: Classes{
				{Id: 3, Address: "上海"},
				{Id: 8, Address: "上海"},
				{Id: 5, Address: "上海"},
			},
		},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
		m[stu.Name].setSexStr()
		for i := range m[stu.Name].Classes {
			m[stu.Name].Classes[i].getClassesName()
		}
	}
	// m = map[string]*student{
	// 	"pprof.cn": &student{Name: "pprof.cn", Age: 18},
	// 	"测试": &student{Name: "测试", Age: 23},
	// 	"博客": &student{Name: "博客", Age: 28},
	// }

	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

	resJson, _ := json.Marshal(map[string](map[string]*student){
		"data": m,
	})
	fmt.Println("resJson: ", string(resJson))

	w.Write([]byte(string(resJson)))
}

type stud struct {
	id   int
	name string
	age  int
}

func demo(ce *[]stud) *[]stud {
	//切片是引用传递，是可以改变值的
	(*ce)[1].age = 999
	*ce = append(*ce, stud{3, "xiaowang", 56})
	return ce
}
func init() {
	ce := []stud{
		{id: 1, name: "zhang", age: 22},
		{2, "wang", 33},
	}
	fmt.Printf("%v\n", ce)
	demo(&ce)
	fmt.Printf("%v\n", ce)

	var a *int = new(int)
	*a = 100

	b := map[string]string{
		"name": "wanglaowu",
	}
	fmt.Println(b, *a)

	var slice = []int{2, 3, 4, 5, 6, 7, 8}

	slice = append(slice, 99)

	fmt.Println(cap(slice), len(slice), slice)
}
