// @author  dreamlu
package models

import (
	"demo/util/log"
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/tool/result"
	time2 "github.com/dreamlu/go-tool/tool/type/time"
	"time"
)

/*user model*/
type User struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	Name       string      `json:"name" valid:"required,len=2-20"`
	Createtime time2.CTime `json:"createtime"` //maybe you like util.JsonDate
}

// dbcrud form data
var db = der.DbCrud{
	Model: User{}, // model
	Table: "user", // table name
}

// get user, by id
func (c *User) GetByID(id string) interface{} {

	var user User // not use *User
	db.ModelData = &user
	if err := db.GetByID(id); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	return result.GetSuccess(user)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearch(params map[string][]string) interface{} {
	var users []*User
	db.ModelData = &users
	pager, err := db.GetBySearch(params)
	if err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetSuccessPager(users, pager)
}

// delete user, by id
func (c *User) Delete(id string) interface{} {

	if err := db.Delete(id); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update user
func (c *User) Update(params map[string][]string) interface{} {

	if err := db.Update(params); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create user
func (c *User) Create(params map[string][]string) interface{} {

	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))

	if err := db.Create(params); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}