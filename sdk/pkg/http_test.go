package pkg

import (
	"fmt"
	"gitlab.benlai.work/go/ymir/sdk/api/response"
	"testing"
)

type ApplicationCicd struct {
	// 应用CicdId
	CicdId int `json:"cicdId"`
	// 应用Id
	AppId int `json:"appId"`
	// 容器内应用类型
	AppType string `json:"appType"`
	// 应用别名
	AliasName string `json:"aliasName"`
	// 应用命名空间
	Namespace string `json:"namespace"`
}

type ApplicationCicdResponse struct {
	response.Response
	Data []ApplicationCicd `json:"data"`
}

func TestGet(t *testing.T) {
	var res ApplicationCicdResponse
	GetJson("http://127.0.0.1:9090/api/v1/appcicd/get/5031", nil, nil, &res)
	fmt.Printf("%v\r\n", res)

	resStr, err := GetString("http://127.0.0.1:9090/api/v1/appcicd/get/5031", nil, nil)
	if err != nil {
		fmt.Sprintf("err: %s", err)
	}
	fmt.Printf("%s\r\n", resStr)
}

func TestPost(t *testing.T) {
	resStr, err := PostString("http://t-jenkins.benlai.cloud/job/project-XX2QnMAPrEwk/job/soa-test/buildWithParameters?DEPLOY_TYPE=完全发布&ENVIRONMENT=branch&IS_AUTO_PUBLISH=True", map[string]string{"Authorization": "Basic YWRtaW46NmhQdy4vKHBbZHhr"}, nil, nil)
	if err != nil {
		fmt.Sprintf("err: %s", err)
	}
	fmt.Printf("%s\r\n", resStr)
}
