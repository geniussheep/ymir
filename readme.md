## 生成swagger文件
- ### 修改 go.mod 
```go
require (
	github.com/gin-gonic/gin v1.7.7
	github.com/spf13/cast v1.5.0
	github.com/swaggo/files v1.0.0
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.4
	gitlab.benlai.work/go/ymir v1.1.8
)
```
- ### 执行 go mod tidy
- ### 执行 swag init --parseDependency --parseInternal
