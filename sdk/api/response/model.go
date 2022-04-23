package response

type Response struct {
	// 请求唯一Id
	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`

	// 业务状态码
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`

	// 返回信息
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`

	//状态
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type response struct {
	Response

	//返回数据
	Data interface{} `json:"data"`
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}
