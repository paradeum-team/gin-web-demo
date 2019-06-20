package router

import (
	"gin-web-demo/bootstrap"
	api "gin-web-demo/controller"
	"github.com/gin-gonic/gin"
	"strings"
	pldconf "gin-web-demo/config"
)

/**
 * 路由配置，并根据配置文件设置根路径
 * 参考url：https://github.com/gin-gonic/gin
 */
func Configure(r *bootstrap.Bootstrapper) {
	prefix :="/"
	//此处可以增加系统应用目录根路径
	pldConf :=pldconf.AppConfig
	contextPath :=pldConf.Server.ContextPath
	if "" !=contextPath && strings.HasPrefix(contextPath,"/"){
		//给拼接好的api ，增加前缀
		prefix=contextPath
	}
	rootRouter :=r.Group(prefix)//设置访问的根目录

	//call concrete route
	concreteRouter(rootRouter)


}

/**
  配置具体的路由信息
 */
func concreteRouter(rootRouter *gin.RouterGroup){
	v1 := rootRouter.Group("/v1")
	v1.GET("/login/:name", api.V1IndexLoginfunc)
	v1.POST("/pwlogin", api.V1PasswordLoginfunc)
	v1.GET("/submit", api.V1IndexSubmitfunc)

	v2 := rootRouter.Group("/v2")
	{//使用花括号，把相关的，组织到一起；
		v2.GET("/submit", api.V2IndexSubmitfunc)
	}

	rootRouter.GET("/v3/submit", api.V2IndexSubmitfunc)

	authorized := rootRouter.Group("/auth")

	authorized.GET("/secret", api.AccessAPIWithAuthorized)

	rootRouter.GET("/exception", api.Exception)
}