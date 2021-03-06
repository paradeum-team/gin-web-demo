package router

import (
	"gin-web-demo/bootstrap"
	pldconf "gin-web-demo/config"
	"gin-web-demo/docs"
	"gin-web-demo/web/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strings"
	"github.com/gin-contrib/pprof"
)

/**
 * 路由配置，并根据配置文件设置根路径
 * 参考url：https://github.com/gin-gonic/gin
 */
func Configure(r *bootstrap.Bootstrapper) {
	prefix := "/"
	//此处可以增加系统应用目录根路径
	pldConf := pldconf.AppConfig
	contextPath := pldConf.Server.ContextPath
	if "" != contextPath && strings.HasPrefix(contextPath, "/") {
		//给拼接好的api ，增加前缀
		prefix = contextPath
	}
	rootRouter := r.Group(prefix) //设置访问的根目录
	pprof.Register(r.Engine)//性能测试工具
	//call concrete route
	concreteRouter(rootRouter)

	// programatically set swagger info
	// programatically set swagger info
	docs.SwaggerInfo.Title = "demo-api:ONLINE API"
	docs.SwaggerInfo.Description = "This is Demo server online restFull api ."
	docs.SwaggerInfo.Version = "1.0"
	/**
	去掉设定的 baseUrl，swagger 也能够顺利的获取到ip，端口
	ipHost :=utils.GetLocalHostIps()[0]
	address := fmt.Sprintf("%s:%d",ipHost, pldConf.Server.Port)
	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.BasePath = prefix
	**/
	rootRouter.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

/**
  配置具体的路由信息
 */
func concreteRouter(rootRouter *gin.RouterGroup) {
	v1 := rootRouter.Group("/v1")
	v1.GET("/users/:name", api.GetUserByName)
	v1.POST("/login", api.Login)
	v1.GET("/users", api.ListUsers)

	v2 := rootRouter.Group("/v2")
	{ //使用花括号，把相关的，组织到一起；
		v2.GET("/submit", api.GetTestUserData)
	}

	ws := rootRouter.Group("/ws")
	{
		ws.GET("/ping",api.Ping)
	}

	accounts := gin.Accounts{
		"admin": "password", //==>{"Authorization":"Basic Zm9vOmJhcg=="}
		"foo":   "bar",      //==>{"Authorization":"Basic YWRtaW46cGFzc3dvcmQ="}
	}
	authorized := rootRouter.Group("/auth", gin.BasicAuth(accounts))
	{
		authorized.GET("/secret", api.AccessAPIWithAuthorized)

	}

	rootRouter.GET("/exception", api.Exception)
}
