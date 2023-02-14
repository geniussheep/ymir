package model

type QueryTest struct {
	Path    string `msgpack:"pathArgs" json:"pathArgs" form:"pathArgs" xml:"pathArgs" uri:"pathArgs" example:"pathArgs"`
	QString string `msgpack:"qString" json:"qString" form:"qString" xml:"qString" uri:"qString" example:"string"`
	QInt    int    `msgpack:"qInt" json:"qInt" form:"qInt" xml:"qInt" uri:"qInt"  binding:"required" example:"123"`
}

type RespTest struct {
	QueryTest
	RespBody string `msgpack:"respBody" json:"respBody" form:"respBody" xml:"respBody" uri:"respBody" example:"string"`
	RespType string `msgpack:"respType" json:"respType" form:"respType" xml:"respType" uri:"respType" example:"string"`
}
