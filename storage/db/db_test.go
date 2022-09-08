package db

import (
	"gitlab.benlai.work/go/ymir/logger"
	dbLogger "gorm.io/gorm/logger"
	"testing"
	"time"
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
	var model KSCSlb
	l := logger.NewLogger(logger.WithLevel(logger.DebugLevel))
	logger.DefaultLogger = l
	l.Log(logger.InfoLevel, "test-logger")
	yorm, err := New(
		SetDsn("sqlserver://plf_user:2KspR9JQw@10.250.60.245:1433?database=BenlaiMonitorNew&connection+timeout=30"),
		SetDriver("mssql"),
		SetLogLevel(dbLogger.LogLevel(logger.DebugLevel.LevelForGorm())))
	if err != nil {
		logger.Errorf("Service.KSCSlb Get error:%e", err.Error())
	}

	where := []interface{}{
		[]interface{}{"environment_type", "=", "trunk"},
		[]interface{}{"curr_used_count", "<", 1022},
	}

	err = yorm.FindByQuery(where, &model)

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

type ApplicationDomainServer struct {

	// 应用domain Id
	Id int `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:domain id"`

	// 服务端应用id
	ServerAppId int `json:"server_app_Id" gorm:"column:server_app_Id;size:256;comment:应用id"`

	// 应用Domain名称
	DomainName string `json:"domain_name" gorm:"column:domain_name;size:256;comment:应用Domain名称"`

	// 应用Domain地址
	DomainAddress string `json:"domain_address" gorm:"column:domain_address;size:256;comment:应用Domain地址"`

	// 应用状态 -- 禁用:disabled, 启用:ready
	Status string `json:"status" gorm:"column:status;size:32;comment:状态"`

	// 服务版本号
	Version string `json:"version" gorm:"column:version;size:46;comment:服务版本号"`

	// 是否是公共服务
	IsPublic DbBool `json:"isPublic" gorm:"column:isPublic;comment:是否是公共服务"`

	// 唯一Id
	SId int `json:"sid" gorm:"column:sid;comment:唯一id"`

	// 是否已删除
	IsDel int `json:"is_del" gorm:"column:is_del;comment:是否已删除"`

	// 创建时间
	CreatedTime *time.Time `gorm:"column:created_time" json:"created_time"`

	// 更新时间
	UpdatedTime *time.Time `gorm:"column:updated_time" json:"updated_time"`

	//应用名称
	AppName string `json:"appName"`

	//应用别名
	AppAliasName string `json:"appAliasName"`
}

func (a *ApplicationDomainServer) Validate() bool {
	return a.Id > 0
}

func (ApplicationDomainServer) TableName() string {
	return "tb_esb_rpc_application_domain_servers"
}

func TestDbmsDb(t *testing.T) {
	var model ApplicationDomainServer
	l := logger.NewLogger(logger.WithLevel(logger.DebugLevel))
	logger.DefaultLogger = l
	l.Log(logger.InfoLevel, "test-logger")
	yorm, err := New(
		SetDsn("root:bb.123#lm.com@tcp(10.250.240.59:3306)/BenlaiEsbBrancheNew?charset=utf8&parseTime=true"),
		SetDriver("mysql"),
		//SetUseDbms(true),
		SetLogLevel(dbLogger.LogLevel(logger.DebugLevel.LevelForGorm())))
	if err != nil {
		logger.Errorf("Service.Esb Get error:%e", err.Error())
	}

	where := []interface{}{
		[]interface{}{"server_app_id", "in", []int{5031, 8338}},
		[]interface{}{"is_del", "=", false},
	}

	err = yorm.FindByQuery(where, &model)

	if err != nil {
		logger.Errorf("Service.Esb Get error:%s", err.Error())
	}

	logger.Infof("get model %v", model)
}
