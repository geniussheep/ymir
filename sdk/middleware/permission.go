package middleware

//
//import (
//	"github.com/casbin/casbin/v2/util"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type UrlInfo struct {
//	Url    string
//	Method string
//}
//
//var casbinExclude = make([]UrlInfo, 0)
//
////AppendCasbinExclude casbin 添加排除权限验证的路由列表
//func AppendCasbinExclude(urlInfos ...UrlInfo) {
//	for _, ui := range urlInfos {
//		casbinExclude = append(casbinExclude, ui)
//	}
//}
//
//// AuthCheckRole 权限检查中间件
//func AuthCheckRole() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		log := api.GetRequestLogger(c)
//		data, _ := c.Get(jwtauth.JwtPayloadKey)
//		v := data.(jwtauth.MapClaims)
//		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
//		var res, isCasbinExclude bool
//		var err error
//		//检查权限
//		if v["rolekey"] == "admin" {
//			res = true
//			c.Next()
//			return
//		}
//		for _, i := range casbinExclude {
//			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
//				isCasbinExclude = true
//				break
//			}
//		}
//		if isCasbinExclude {
//			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
//			c.Next()
//			return
//		}
//		res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
//		if err != nil {
//			log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
//			response.Error(c, 500, err, "")
//			return
//		}
//
//		if res {
//			log.Infof("isTrue: %v role: %s method: %s path: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path)
//			c.Next()
//		} else {
//			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
//			c.JSON(http.StatusOK, gin.H{
//				"code": 403,
//				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
//			})
//			c.Abort()
//			return
//		}
//
//	}
//}
