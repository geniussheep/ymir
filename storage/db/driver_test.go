package db

import "testing"

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
