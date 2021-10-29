package db

import (
	"errors"
	"gitlab.benlai.work/go/dbms"
	"log"
	"os"
	"time"

	// mysql driver
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Yorm interface
type Yorm interface {
	Init(models ...interface{}) error
	FindOne(id int64, model interface{}) error
	First(model interface{})
	Last(model interface{})
	FindAll(models interface{}) error
	FindByQuery(where interface{}, models interface{}) error
	FindByQueryForPage(wheres interface{}, columns interface{}, orderBy interface{}, page, rows int, models interface{}) error
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Db() *gorm.DB
}

type yorm struct {
	dsn    string
	driver Driver
	db     *gorm.DB
}

func NewWithDbms(connStr string) (Yorm, error) {
	dbms := dbms.NewClient()
	driver, dsn, err := dbms.GetConnectionString(connStr)
	if err != nil {
		return nil, err
	}
	return New(dsn, Driver(driver))
}

// NewYorm func
func New(dsn string, driver Driver) (Yorm, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	r := &yorm{
		dsn:    dsn,
		driver: driver,
		db:     nil,
	}
	switch r.driver {
	case MYSQL:
		db, err := gorm.Open(mysql.Open(r.dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		r.db = db
		break
	case MSSQL:
		db, err := gorm.Open(sqlserver.Open(r.dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		r.db = db
		break
	}

	return r, nil

}

func (r *yorm) Init(models ...interface{}) error {
	if models == nil || len(models) == 0 {
		return errors.New("model is missing")
	}

	if err := r.db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}

func (r *yorm) FindOne(id int64, model interface{}) error {
	if id <= 0 {
		return errors.New("id is invalid")
	}

	r.db.First(model, id)

	return nil
}

func (r *yorm) First(model interface{}) {
	r.db.First(model)
}

func (r *yorm) Last(model interface{}) {
	r.db.Last(model)
}

func (r *yorm) FindAll(models interface{}) error {

	if err := r.db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) FindByQuery(where interface{}, models interface{}) error {

	db, err := BuildWhere(r.db, where)
	if err != nil {
		return err
	}

	if err := db.Find(&models).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) FindByQueryForPage(where interface{}, columns interface{}, orderBy interface{}, page, rows int, models interface{}) error {

	db, err := BuildWhereForPage(r.db, where, columns, orderBy, page, rows)
	if err != nil {
		return err
	}

	if err := db.Find(&models).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) Create(model interface{}) error {
	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) Update(model interface{}) error {
	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) Delete(model interface{}) error {

	if err := r.db.Delete(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *yorm) Db() *gorm.DB {
	return r.db
}
