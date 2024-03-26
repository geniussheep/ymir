package db

import (
	"github.com/geniussheep/ymir/logger"
	dbLogger "gorm.io/gorm/logger"
	"testing"
)

func TestDriver(t *testing.T) {
	//d := Driver("mssql")
	d := Driver("sqlserver")
	switch d {
	case MYSQL:
		println("mysql-test")
		break
	case MSSQL, SQLSERVER:
		println("mssql-test")
		break
	default:
		println("default -- test")
		break
	}
}

type Application struct {
	Id                    int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name                  string `json:"name" gorm:"size:256;comment:handle"`
	ApplicationType       string `json:"application_type" gorm:"size:256;comment:应用类型"`
	Description           string `json:"description" gorm:"size:512;comment:描述	"`
	ApplicationKey        string `json:"application_key" gorm:"size:256;comment:key"`
	ApplicationCategoryId int    `json:"application_category_id" gorm:"comment:应用类型"`
	Status                string `json:"status" gorm:"size:32;comment:key"`
	AliasName             string `json:"alias_name" gorm:"size:256;comment:key"`
	IsHasDocker           bool   `json:"is_has_docker" gorm:"comment:key"`
}

func (Application) TableName() string {
	return "tb_bel_application"
}

// ViewApplicationContact 应用联系人
type ViewApplicationContact struct {
	// 应用Id
	ApplicationId int `json:"applicationId" gorm:"column:application_id;comment:应用id"`

	// 应用名称
	ApplicationName string `json:"applicationName" gorm:"column:application_name;size:256;comment:handle"`

	// 应用Key
	ApplicationKey string `json:"applicationKey" gorm:"column:application_key;size:256;comment:应用key"`

	// 应用是否容器化
	IsHasDocker bool `json:"isHasDocker" gorm:"column:is_has_docker;comment:是否容器化"`

	// 应用别名
	AliasName string `json:"aliasName" gorm:"column:alias_name;size:256;comment:应用别名"`

	// 应用状态 -- 禁用:disabled, 启用:ready, 待审核:unready, 重用:reused
	Status string `json:"status" gorm:"column:status;size:32;comment:应用状态"`

	// 联系人Id
	ContactId int `json:"contactId" gorm:"column:contact_id;comment:联系人Id"`

	// 联系人姓名
	ContactName string `json:"contactName" gorm:"column:contact_name;size:256;comment:联系人姓名"`

	// 联系人邮件
	Email string `json:"email" gorm:"column:email;size:256;comment:联系人邮件"`

	// 联系人手机
	Sms string `json:"sms" gorm:"column:sms;size:256;comment:联系人手机"`

	// 联系人微信
	WeChat string `json:"wechat" gorm:"column:wechat;size:256;comment:联系人微信"`

	// 联系人系统Id
	SysUserId int `json:"sysUserId" gorm:"column:sys_user_id;comment:联系人系统Id"`
}

func (a *ViewApplicationContact) Validate() bool {
	return a.ApplicationId > 0 && a.ContactId > 0
}

func (ViewApplicationContact) TableName() string {
	return "view_bel_application_contact"
}

type KSCSlb struct {
	// 服务器Id
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:主键编码"`

	// 环境
	EnvironmentType string `json:"environment_type" gorm:"column:environment_type;size:256;comment:环境"`

	// 金山云的slb类型
	Type string `json:"type" gorm:"column:type;comment:金山云的slb类型"`

	// 金山云slb网段id
	SubnetId string `json:"subnetid" gorm:"column:subnetid;comment:金山云slb网段id"`

	// 当前网段使用数量
	CurrUsedCount int `json:"curr_used_count" gorm:"column:curr_used_count;comment:当前网段使用数量"`
}

func (KSCSlb) TableName() string {
	return "tb_bel_ksc_slb"
}

func TestDb(t *testing.T) {
	var models []Application
	var model KSCSlb
	l := logger.NewLogger(logger.WithLevel(logger.DebugLevel))
	logger.DefaultLogger = l
	l.Log(logger.InfoLevel, "test-logger")
	yorm, err := New(
		SetDsn("sqlserver://plf_user:2KspR9JQw@10.250.60.245:1433?database=BenlaiMonitorNew&connection+timeout=30"),
		SetDriver("mssql"),
		SetLogLevel(dbLogger.LogLevel(logger.DebugLevel.LevelForGorm())))
	if err != nil {
		logger.Errorf("open sql connecting error:%e", err.Error())
	}
	//appIds := []int{1, 2, 4, 5031}
	//appNames := []string{"api", "risk", "soa-test"}
	//where := []interface{}{
	//	[]interface{}{"application_id", "in", appIds},
	//	[]interface{}{"application_name in ? or alias_name in ?", appNames, appNames},
	//}

	where := []interface{}{
		[]interface{}{"status = ?", "ready"},
	}

	err = yorm.FindByQueryForPage(where, "id desc", 20, 10, &models)

	err = yorm.FindByQueryForPage(where, "id desc", 4, 20, &models)
	return

	if err != nil {
		logger.Errorf("Service.KSCSlb Get error:%s", err.Error())
	}
	updateFileds := map[string]interface{}{"curr_used_count": 12}

	where = []interface{}{
		[]interface{}{"id", "=", model.Id},
	}
	err = yorm.UpdateBatch(updateFileds, where, &model)

	err = yorm.FindByQuery(where, &model)

	logger.Infof("get model %v", model)
}
