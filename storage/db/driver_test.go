package db

import "testing"

func TestDriver(t *testing.T) {
	d := Driver("mssql")
	switch d {
	case MYSQL:
		println("mysql-test")
		break
	case MSSQL:
	case SQLSERVER:
		println("mssql-test")
		break
		println("sql-test")
		break
	default:
		println("default -- test")
		break
	}
}
