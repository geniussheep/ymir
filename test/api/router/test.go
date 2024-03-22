package router

import (
	"github.com/geniussheep/ymir/sdk"
	ymirApi "github.com/geniussheep/ymir/sdk/api"
	testApi "github.com/geniussheep/ymir/test/api"
)

func init() {
	ApiRouterScan = append(ApiRouterScan, func() {
		api := testApi.Test{}
		sdk.Runtime.GetWebApi().AppendRouters(V1, ymirApi.RouterEntry{Path: "test/:pathArgs/get", Method: "GET", Handler: api.Get})
		sdk.Runtime.GetWebApi().AppendRouters(V1, ymirApi.RouterEntry{Path: "app/get/:appId/:environment", Method: "GET", Handler: api.GetApplication})
		sdk.Runtime.GetWebApi().AppendRouters(V1, ymirApi.RouterEntry{Path: "app/id/:appId/build/create", Method: "POST", Handler: api.CreateByAppId})
	})
}
