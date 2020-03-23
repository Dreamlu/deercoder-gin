package client

import (
	"demo/models/client"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/validator"
	"github.com/dreamlu/gt/tool/xss"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var p client.Client

//根据id获得data
func GetByID(u *gin.Context) {
	id := u.Query("id")
	ss := p.GetByID(id)
	u.JSON(http.StatusOK, ss)
}

//data信息分页
func GetBySearch(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	ss := p.GetBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//data信息删除
func Delete(u *gin.Context) {
	id := u.Param("id")
	ss := p.Delete(id)
	u.JSON(http.StatusOK, ss)
}

//data信息修改
func Update(u *gin.Context) {
	var data client.Client

	// json 类型需要匹配
	// 与spring boot不同
	// 不能自动将字符串转成对应类型
	// 严格匹配
	_ = u.ShouldBindJSON(&data)
	// do something

	ss := p.Update(&data)
	u.JSON(http.StatusOK, ss)
}

//新增data信息
func Create(u *gin.Context) {
	var data client.Client

	// 自定义日期格式问题
	err := u.ShouldBindJSON(&data)
	log.Println(err)

	ss := p.Create(&data)
	u.JSON(http.StatusOK, ss)
}

//data信息修改
func UpdateForm(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)

	ss := p.UpdateForm(values)
	u.JSON(http.StatusOK, ss)
}

//新增data信息
func CreateForm(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)                              //html特殊字符转换
	res := validator.Valid(values, client.Client{}) //验证规则
	if res != result.MapValSuccess {
		u.JSON(http.StatusOK, res)
		return
	}

	ss := p.CreateForm(values)
	u.JSON(http.StatusOK, ss)
}