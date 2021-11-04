package db

import "testing"

func TestDriver(t *testing.T) {
	d := Driver("mssql")
	switch d {
	case MYSQL:
		println("mysql-test")
		break
	case MSSQL:
		println("mssql-test")
		break
	default:
		println("default -- test")
		break
	}
}
