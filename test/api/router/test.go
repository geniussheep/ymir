package router

import (
	ymirApi "gitlab.benlai.work/go/ymir/sdk/api"
	testApi "gitlab.benlai.work/go/ymir/test/api"
)

func init() {
	ApiRouterScan = append(ApiRouterScan, func() {
		api := testApi.Test{}
		TestApi.AppendRouters(V1, ymirApi.RouterEntry{Path: "test/:pathArgs/get", Method: "GET", Handler: api.Get})
		TestApi.AppendRouters(V1, ymirApi.RouterEntry{Path: "app/get/:appId", Method: "GET", Handler: api.GetApplication})
	})
}
