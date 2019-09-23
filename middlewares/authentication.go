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
		router := c.ResourcePath
		uri := c.Request.RequestURI
		method := c.Request.Method
		fmt.Printf("authorization=%v \n", authorization)

		fmt.Println("uri=" + uri)
		fmt.Println("router...=" + router)
		fmt.Println("method=" + method)

		hasPermission := false
		//把不需要验证的，都过滤掉。
		if strings.Contains(router, "/api/swagger") || !strings.HasPrefix(router, "/dsp") || strings.Contains(router, "/v1/login") {
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
