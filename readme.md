# ymir 用于go的webapi基础框架

## 使用ymir框架生成swagger文件步骤
- ### 修改 go.mod 
```go
require (
	github.com/gin-gonic/gin v1.8.1
	github.com/spf13/cast v1.5.0
	github.com/swaggo/files v1.0.0
	github.com/swaggo/gin-swagger v1.5.3
	github.com/swaggo/swag v1.8.10
	gitlab.benlai.work/go/ymir v1.1.8
)
```
- ### 执行 go mod tidy
- ### 执行 swag init --parseDependency --parseInternal
