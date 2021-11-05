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

func TestDb(t *testing.T) {
	var model Application
	yorm, err := New(
		SetDsn("sqlserver://plf_user:2KspR9JQw@192.168.60.245:1433?database=BenlaiMonitorNew&connection+timeout=30"),
		SetDriver("mssql"),
		SetLogLevel(dbLogger.Info))
	if err != nil {
		logger.Errorf("Service.Application Get error:%e", err.Error())
	}
	err = yorm.FindOne(32, &model)
	if err != nil {
		logger.Errorf("Service.Application Get error:%s", err.Error())
	}
	logger.Infof("get model %v", model)
}
