# ymir 用于go的webapi基础框架

## 使用ymir框架生成swagger文件步骤
- ### 安装 swag cmd : go install github.com/swaggo/swag/cmd/swag@latest
- ### 执行 go mod tidy
- ### 执行 swag init --parseDependency --parseInternal

## 版本更新列表
### v1.0.0
- First stable version.
### v1.0.1
- Api的response功能更新：
    - 新增OKWithCustomCode方法：通常成功数据处理且自定义response.code
    - 调整方法名拼写错误：Custum -> Custom
- Api功能更新：
    -  新增OKWithCustomCode方法：通常成功数据处理且自定义response.code
### v1.0.2
- 调整数据库方法及注释
    - 去除FindByQueryForPage方法的column参数
    - 新增yorm内相关方法的注释,主要解释where条件语句如何code
### v1.0.3
- 调整数据库方法及注释
    - FindByQueryForPage方法新增total参数
### v1.0.4
- Api返回值Response
  - 去除每个Response字段json设置的,omitempty，保证不管什么值都会返回不会被丢弃(当字段是0或nil时)
### v1.0.5
- Http方法优化
  - 加入默认的请求头 "Content-Type": "application/json"
- validator方法优化
  - 调整 CheckArrayArgs 入参为 value []interface{}
### v1.0.6
- Yorm数据库调整
  - 调整日志格式
### v1.0.6.1
- 升级viper版本修复bug
- 去除test代码
### v1.0.7
- 新增http相关配置
- 新增中间件 ResponseHeader.go
- 调整跨域头的逻辑支持配置
- 调整pkg.CheckArrayArgs方法支持泛型