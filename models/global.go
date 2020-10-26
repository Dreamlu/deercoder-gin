package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/type/time"
)

type IDCom struct {
	ID uint64 `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
}

// 通用模型
type ModelCom struct {
	IDCom
	CreateTime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间自动生成
}

// 多账号, 如果不需要多账号, 注释掉AdminID即可(ps: 为了简化部署,可直接在util/db/db.go中加入初始化appid applet账号信息)
// 账号关联
type AdminCom struct {
	ModelCom
	AdminID uint64 `json:"admin_id" gorm:"type:bigint(20);INDEX:查询索引admin_id"`
}

// ================ common ============
type DN interface {
	New() DN
}

// common crud
type Com struct {
	Model     interface{}
	Data      DN
	ArrayData interface{}
}

var crud = gt.NewCrud()

// get data, by id
func (c *Com) Get(params cmap.CMap) (data interface{}, err error) {
	crud.Params(gt.Model(c.Model), gt.Data(&c.Data))
	if err = crud.GetByData(params).Error(); err != nil {
		return
	}
	return c.Data, nil
}

// get data, limit and search
func (c *Com) Search(params cmap.CMap) (datas interface{}, pager result.Pager, err error) {
	crud.Params(gt.Model(c.Model), gt.Data(&c.ArrayData))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		return nil, pager, cd.Error()
	}
	return datas, cd.Pager(), nil
}

// delete data, by id
func (c *Com) Delete(id interface{}) error {

	return crud.Params(gt.Model(c.Model)).Delete(id).Error()
}

// update data
func (c *Com) Update() error {

	crud.Params(gt.Data(c.Data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	return nil
}

// create data
func (c *Com) Create() error {

	crud.Params(gt.Data(c.Data))
	if err := crud.Create().Error(); err != nil {
		return err
	}
	return nil
}
