package db

import (
	"gitlab.benlai.work/go/ymir/logger"
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
	err = yorm.UpdateBatch(updateFileds, where)

	err = yorm.FindByQuery(where, &model)

	logger.Infof("get model %v", model)
}
