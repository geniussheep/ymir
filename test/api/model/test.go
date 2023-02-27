package model

type QueryTest struct {
	Path    string `msgpack:"pathArgs" json:"pathArgs" form:"pathArgs" xml:"pathArgs" uri:"pathArgs" example:"pathArgs"`
	QString string `msgpack:"qString" json:"qString" form:"qString" xml:"qString" uri:"qString" example:"string"`
	QInt    int    `msgpack:"qInt" json:"qInt" form:"qInt" xml:"qInt" uri:"qInt"  binding:"required" example:"123"`
}

type RespTest struct {
	QueryTest
	RespBody string `msgpack:"respBody" json:"respBody" form:"respBody" xml:"respBody" uri:"respBody" example:"string"`
	RespType string `msgpack:"respType" json:"respType" form:"respType" xml:"respType" uri:"respType" example:"string"`
}

type Application struct {
	// 应用Id
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:主键编码"`

	// 应用名称
	Name string `json:"name" gorm:"column:name;size:256;comment:handle"`

	// 应用类型
	ApplicationType string `json:"applicationType" gorm:"column:application_type;size:256;comment:应用类型"`

	// 应用描述
	Description string `json:"description" gorm:"column:description;size:512;comment:描述	"`

	// 应用Key
	ApplicationKey string `json:"applicationKey" gorm:"column:application_key;size:256;comment:应用key"`

	// 应用类型
	ApplicationCategoryId int `json:"applicationCategoryId" gorm:"column:application_category_id;comment:应用类型"`

	// 应用状态 -- 禁用:disabled, 启用:ready, 待审核:unready, 重用:reused
	Status string `json:"status" gorm:"column:status;size:32;comment:状态"`

	// 应用别名
	AliasName string `json:"aliasName" gorm:"column:alias_name;size:256;comment:应用别名"`

	// 应用是否容器化
	IsHasDocker bool `json:"isHasDocker" gorm:"column:is_has_docker;comment:是否容器化"`

	// 应用是否以上云
	IsInKsc bool `json:"isInKsc" gorm:"column:is_in_ksc;comment:是否已上云"`
}

func (a *Application) Validate() bool {
	return a.Id > 0
}

func (Application) TableName() string {
	return "tb_bel_application"
}
