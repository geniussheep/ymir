package pkg

type (
	Mode string
)

const (
	Dev  Mode = "dev"  //开发模式
	Test Mode = "test" //测试模式
	Prod Mode = "prod" //生产模式
)

func (e Mode) String() string {
	return string(e)
}
