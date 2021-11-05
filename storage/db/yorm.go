package db

import (
	"context"
	"errors"
	"fmt"
	"gitlab.benlai.work/go/dbms"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// driver
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

// Yorm interface
type Orm interface {
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
	WithContext(ctx context.Context) *Yorm
	Db() *gorm.DB
}

type Yorm struct {
	dsn    string
	driver Driver
	db     *gorm.DB
}

// NewYorm func
func New(opts ...Option) (*Yorm, error) {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		op.logConfig,
	)

	if len(op.Dsn) <= 0 {
		return nil, fmt.Errorf("db args:Dsn is empty")
	}

	if !op.UseDbms && len(op.Driver) <= 0 {
		return nil, fmt.Errorf("db args:Driver is empty")
	}

	if op.UseDbms {
		dbmsClient := dbms.NewClient()
		driver, dsn, err := dbmsClient.GetConnectionString(op.Dsn)
		op.Dsn = dsn
		op.Driver = driver
		if err != nil {
			return nil, err
		}
	}

	r := &Yorm{
		dsn:    op.Dsn,
		driver: Driver(op.Driver),
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
	case MSSQL, SQLSERVER:
		db, err := gorm.Open(sqlserver.Open(r.dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		r.db = db
		break
	case PGSQL:
		db, err := gorm.Open(postgres.Open(r.dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		r.db = db
		break
	case SQLITE:
		db, err := gorm.Open(sqlite.Open(r.dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		r.db = db
		break
	default:
		return nil, fmt.Errorf("the db driver: %s unknow", op.Driver)
	}
	return r, nil
}

func (r *Yorm) Init(models ...interface{}) error {
	if models == nil || len(models) == 0 {
		return errors.New("model is missing")
	}

	if err := r.db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}

func (r *Yorm) FindOne(id int64, model interface{}) error {
	if id <= 0 {
		return errors.New("id is invalid")
	}

	r.db.First(model, id)

	return nil
}

func (r *Yorm) First(model interface{}) {
	r.db.First(model)
}

func (r *Yorm) Last(model interface{}) {
	r.db.Last(model)
}

func (r *Yorm) FindAll(models interface{}) error {

	if err := r.db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) FindByQuery(where interface{}, models interface{}) error {

	db, err := BuildWhere(r.db, where)
	if err != nil {
		return err
	}

	if err := db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) FindByQueryForPage(where interface{}, columns interface{}, orderBy interface{}, page, rows int, models interface{}) error {

	db, err := BuildWhereForPage(r.db, where, columns, orderBy, page, rows)
	if err != nil {
		return err
	}

	if err := db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Create(model interface{}) error {
	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Update(model interface{}) error {
	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Delete(model interface{}) error {

	if err := r.db.Delete(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Db() *gorm.DB {
	return r.db
}

func (r *Yorm) WithContext(ctx context.Context) *Yorm {
	r.db.WithContext(ctx)
	return r
}
