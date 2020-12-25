package middlewares

import (
	"fmt"
	"gin-web-demo/common/app"
	"gin-web-demo/common/e"
	"github.com/gin-gonic/gin"
	"go/types"
	"net/http"
	"strings"
)

/**
 * http api 鉴权
 */
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("authorization")
		//router := c.ResourcePath //1.4.2 自定义的gin 版本
		uri := c.Request.RequestURI
		url := c.Request.URL.Path
		for _, p := range c.Params {//变种替换，不依赖 版本
			url = strings.Replace(url, p.Value, fmt.Sprintf(":%s", p.Key), 1)
		}
		router := url
		method := c.Request.Method
		fmt.Printf("authorization=%v \n", authorization)

		fmt.Println("uri=" + uri)
		fmt.Println("router...=" + router)
		fmt.Println("method=" + method)

		hasPermission := false
		//把不需要验证的，都过滤掉。
		if strings.Contains(uri, "/api/") || strings.Contains(router, "/api:any") || !strings.HasPrefix(router, "/dsp") || strings.Contains(router, "/v1/login") || strings.Contains(router, "/ws/ping") {
			hasPermission = true
			c.Next()
			return
		}

		//toDo authentication 逻辑---token 是否有效
		//user,e := serviceauth.NewAuthService().CheckTokenExpiration(authorization)
		//if e!=nil {
		//	c.Abort()
		//	app.NewResponse(c,http.StatusUnauthorized,gin.H{"message":string(e.Error())})
		//}else {
		//toDo authentication 逻辑---token 是否有权访问

		//	hasPermission,err:=authority.NewInstance().HasPermission(*user,method,router)
		//	if err !=nil{
		//		c.Abort()
		//		app.NewResponse(c,http.StatusUnauthorized,gin.H{"message":string(e.Error())})
		//	}else if !hasPermission{
		//		c.Abort()
		//		app.NewResponse(c,http.StatusUnauthorized,gin.H{"message":"没有访问权限"})
		//	}

		//c.Next() //继续访问
		//}

		if hasPermission {
			c.Next() //继续访问
		} else {
			c.Abort()
			app.NewBadResponse(c, http.StatusUnauthorized, e.ERR_NO_AUTH, types.Nil{})
		}
	}
}
