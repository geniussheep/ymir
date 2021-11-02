module gitlab.benlai.work/go/ymir/sdk

go 1.16

require (
	github.com/casbin/casbin/v2 v2.38.1
	github.com/gin-gonic/gin v1.7.4
	github.com/google/uuid v1.3.0
	gitlab.benlai.work/go/ymir v1.0.2
	gitlab.benlai.work/go/ymir/plugins/logger/zap v1.0.0
	gorm.io/gorm v1.22.0
)

replace (
	gitlab.benlai.work/go/ymir => ../
	gitlab.benlai.work/go/ymir/plugins/logger/zap => ../plugins/logger/zap
)
