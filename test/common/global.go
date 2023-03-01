package common

import "fmt"

const (
	// Version go-admin version info
	Version   = "1.0.0"
	DbMonitor = "dbMonitor"
	TestRedis = "testRedis"
	AppName   = "TestCmd"
	Desc      = "测试Ymir"
	Zk        = "%s-zk"
)

func GetZkName(env string) string {
	return fmt.Sprintf(Zk, env)
}
