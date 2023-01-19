package db

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// BuildWhere 构建where条件
//
//1、and条件测试
//where := []interface{}{
//	[]interface{}{"id", "=", 1},
//	[]interface{}{"username", "chen"},
//}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (id = 1)and(username = 'chen')
//
//2、结构体条件测试
//where := user.User{ID: 1, UserName: "chen"}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (id = 1) and (username = 'chen')
//
//3、in,or条件测试
//where := []interface{}{
//	[]interface{}{"id", "in", []int{1, 2}},
//	[]interface{}{"username", "=", "chen", "or"},
//}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (id in ('1','2')) OR (username = 'chen')
//
//3.1、not in,or条件测试
//where := []interface{}{
//	[]interface{}{"id", "not in", []int{1}},
//	[]interface{}{"username", "=", "chen", "or"},
//}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (id not in ('1')) OR (username = 'chen')
//
//4、map条件测试
//where := map[string]interface{}{"id": 1, "username": "chen"}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (`users`.`id` = '1') AND (`users`.`username` = 'chen')
//
//5、and,or混合条件测试
//where := []interface{}{
//	[]interface{}{"id", "in", []int{1, 2}},
//	[]interface{}{"username = ? or nickname = ?", "chen", "yond"},
//}
//db, err = entity.BuildWhere(db, where)
//db.Find(&users)
//// SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen' or nickname = 'yond')
//
////注：不要使用下方方法
///*
//where := []interface{}{
//	[]interface{}{"id", "in", []int{1, 2}},
//	[]interface{}{
//		[]interface{}{"username", "=", "chen"},
//		[]interface{}{"username", "=", "yond", "or"},
//	},
//}
//// 返回sql: SELECT * FROM `users`  WHERE (id in ('1','2')) AND (username = 'chen') OR (username = 'yond')
//// 与设想不一样
//// 经过测试，gorm底层暂时不支持db.Where(func(db *gorm.DB) *gorm.DB {})闭包方法
//*/
func BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	var err error
	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					return nil, errors.New("切片长度不能小于2")
				}
				columnstr := column.(string)
				// 拼接参数形式
				if strings.Index(columnstr, "?") > -1 {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/

					if strings.Index(" in notin ", _opt) > -1 {
						// val 是数组类型
						column = columnstr + " " + opt + " (?)"
					} else if strings.Index(" = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp", _opt) > -1 {
						column = columnstr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			} else if t == reflect.Map /*Map*/ {
				db = db.Where(item)
			} else {
				/*
					// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						db, err = BuildWhere(db, item)
						if err != nil {
							panic(err)
						}
						return db
					})*/

				db, err = BuildWhere(db, item)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		return nil, errors.New("参数有误")
	}
	return db, nil
}

func BuildWhereForPage(db *gorm.DB, where interface{}, columns interface{}, orderBy interface{}, page, rows int) (*gorm.DB, error) {
	var err error
	db, err = BuildWhere(db, where)
	if err != nil {
		return nil, err
	}
	db = db.Select(columns)
	if orderBy != nil && orderBy != "" {
		db = db.Order(orderBy)
	}
	if page > 0 && rows > 0 {
		db = db.Limit(rows).Offset((page - 1) * rows)
	}
	return db, err
}
