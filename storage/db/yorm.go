package db

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	// driver
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

// Orm -- Yorm interface
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

// New Yorm func
func New(opts ...Option) (*Yorm, error) {
	op := setDefault()
	for _, o := range opts {
		if o != nil {
			o(&op)
		}
	}

	if len(op.Dsn) <= 0 {
		return nil, fmt.Errorf("db args:Dsn is empty")
	}

	r := &Yorm{
		dsn:    op.Dsn,
		driver: Driver(op.Driver),
		db:     nil,
	}

	newLogger := NewYormLogger(op.logConfig)

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

func (r *Yorm) checkIsInit() error {
	if r == nil || r.db == nil {
		return errors.New("yorm is not init，please check code")
	}
	return nil
}

func (r *Yorm) Init(models ...interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if models == nil || len(models) == 0 {
		return errors.New("model is missing")
	}

	if err := r.db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}

func (r *Yorm) FindOne(id int64, model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if id <= 0 {
		return errors.New("id is invalid")
	}

	r.db.First(model, id)

	return nil
}

func (r *Yorm) First(model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}
	r.db.First(model)
	return nil
}

func (r *Yorm) Last(model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}
	r.db.Last(model)
	return nil
}

func (r *Yorm) FindAll(models interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}
	if err := r.db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

// FindByQuery 根据条件查询数据
// where: 构建where条件
//
//	1、and条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "=", 1},
//			[]interface{}{"username", "chen"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id = 1)and(username = 'chen')
//
//	2、结构体条件测试
//	where := user.User{ID: 1, UserName: "chen"}
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id = 1) and (username = 'chen')
//
//	3、in,or条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{"username", "=", "chen", "or"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id in ('1','2')) OR (username = 'chen')
//
//	3.1、not in,or条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "not in", []int{1}},
//			[]interface{}{"username", "=", "chen", "or"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id not in ('1')) OR (username = 'chen')
//
//	4、map条件测试
//	where := map[string]interface{}{"id": 1, "username": "chen"}
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (`users`.`id` = '1') AND (`users`.`username` = 'chen')
//
//	5、and,or混合条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{"username = ? or nickname = ?", "chen", "yond"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen' or nickname = 'yond')
//
//	注：不要使用下方方法
//	where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{
//				[]interface{}{"username", "=", "chen"},
//				[]interface{}{"username", "=", "yond", "or"},
//			},
//		}
//
//	返回sql: SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen') OR (username = 'yond')
//	与设想不一样
//	经过测试，gorm底层暂时不支持db.Where(func(db *gorm.DB) *gorm.DB {})闭包方法
func (r *Yorm) FindByQuery(where interface{}, models interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}
	db, err := BuildWhere(r.db, where)
	if err != nil {
		return err
	}

	if err := db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

// FindByQueryForPage  根据条件分页查询数据
// where: 构建where条件
//
//	1、and条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "=", 1},
//			[]interface{}{"username", "chen"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id = 1)and(username = 'chen')
//
//	2、结构体条件测试
//	where := user.User{ID: 1, UserName: "chen"}
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id = 1) and (username = 'chen')
//
//	3、in,or条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{"username", "=", "chen", "or"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id in ('1','2')) OR (username = 'chen')
//
//	3.1、not in,or条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "not in", []int{1}},
//			[]interface{}{"username", "=", "chen", "or"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (id not in ('1')) OR (username = 'chen')
//
//	4、map条件测试
//	where := map[string]interface{}{"id": 1, "username": "chen"}
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	// SELECT * FROM `users`  WHERE (`users`.`id` = '1') AND (`users`.`username` = 'chen')
//
//	5、and,or混合条件测试
//
//		where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{"username = ? or nickname = ?", "chen", "yond"},
//		}
//
//	db, err = entity.BuildWhere(db, where)
//	db.Find(&users)
//	SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen' or nickname = 'yond')
//
//	注：不要使用下方方法
//	where := []interface{}{
//			[]interface{}{"id", "in", []int{1, 2}},
//			[]interface{}{
//				[]interface{}{"username", "=", "chen"},
//				[]interface{}{"username", "=", "yond", "or"},
//			},
//		}
//
//	返回sql: SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen') OR (username = 'yond')
//	与设想不一样
//	经过测试，gorm底层暂时不支持db.Where(func(db *gorm.DB) *gorm.DB {})闭包方法
//
// orderBy: 排序字段 例： interface{}{“id desc”}
// pageIndex: 当前第几页
// pageSize:  每页数据条数
// total:  数据总条数
func (r *Yorm) FindByQueryForPage(where interface{}, orderBy interface{}, pageIndex, pageSize int, total *int64, models interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	db, err := BuildWhere(r.db, where)
	if err != nil {
		return err
	}

	if orderBy != nil && orderBy != "" {
		db = db.Order(orderBy)
	}
	if pageIndex > 0 && pageSize > 0 {
		db = db.Model(models).Count(total).Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	}

	if err := db.Find(models).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Create(model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Update(model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if model == nil {
		return errors.New("model is missing")
	}

	if err := r.db.Save(model).Error; err != nil {
		return err
	}

	return nil
}

// UpdateBatch 批量更新 支持多个字段更新
// updateFileds:待更新的字段
//
//		例1、使用map方式传入 - map[string]interface{}{"name": "hello", "age": 18, "active": false}
//	 	例2、使用struct方式传入 - User{Name: "hello", Age: 18, Active: false}
//
// where: 查询条件
func (r *Yorm) UpdateBatch(updateFileds interface{}, where interface{}, model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if updateFileds == nil {
		return errors.New("updateFileds is missing")
	}

	if where == nil {
		return errors.New("where condition is missing")
	}

	db, err := BuildWhere(r.db, where)
	if err != nil {
		return err
	}

	if err := db.Model(&model).Updates(updateFileds).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Delete(model interface{}) error {
	if err := r.checkIsInit(); err != nil {
		return err
	}

	if err := r.db.Delete(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *Yorm) Db() *gorm.DB {
	return r.db
}

func (r *Yorm) WithContext(ctx context.Context) *Yorm {
	r.db = r.db.WithContext(ctx)
	return r
}
