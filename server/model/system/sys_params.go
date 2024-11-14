package system

import "github.com/SevenCryber/my-admin/server/global"

type SysParams struct {
	global.YAC_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;comment:参数名称;" binding:"required"`   //参数名称
	Key   string `json:"key" form:"key" gorm:"column:key;comment:参数键;" binding:"required"`       //参数键
	Value string `json:"value" form:"value" gorm:"column:value;comment:参数值;" binding:"required"` //参数值
	Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:参数说明;"`                      //参数说明
}

// TableName 参数 SysParams自定义表名 sys_params
func (SysParams) TableName() string {
	return "sys_params"
}