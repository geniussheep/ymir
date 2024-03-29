package db

type Driver string

const (
	MYSQL     Driver = "mysql"
	MSSQL     Driver = "mssql"
	SQLSERVER Driver = "sqlserver"
	PGSQL     Driver = "pgsql"
	SQLITE    Driver = "sqlite"
)
