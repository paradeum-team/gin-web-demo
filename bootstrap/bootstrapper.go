package bootstrap

import (
	"fmt"
	"gin-web-demo/common/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*gin.Engine
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Engine:       gin.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//gin.SetMode(gin.ReleaseMode) //系统启动日志输出-- release
	gin.SetMode(gin.DebugMode)   //系统启动日志输出--debug
	if "release" !=gin.Mode(){//如果不是release 的时候，就启用日志中间件
		//设置日志级别或者中间件
		b.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("[%s] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				b.AppName,
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
	}
	b.Use(gin.Recovery())
	plogger.NewInstance().GetLogger().SetLevel("info")//业务日志级别--可用于controller，service，dao 等

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...Configurator) {
	b.Run(addr)
}
