package controller

import (
	"gin-web-demo/common/utils"
	"gin-web-demo/service"
	"gin-web-demo/vo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func V1PasswordLoginfunc(c *gin.Context) {
	plogger.NewInstance().GetLogger().Info("user  login, will access ...")

	var json vo.LoginJSON
	// If EnsureBody returns false, it will write automatically the error
	// in the HTTP stream and return a 400 error. If you want custom error
	// handling you should use: c.ParseBody(interface{}) error
	log.Println("come here 123")

	if err:=c.ShouldBindJSON(&json);err==nil {
		if json.User == "pld" && json.Password == "pld" {
			c.JSON(200, gin.H{"status": "you are logged in","ext":"ext345"})
		} else {
			//data :=map[string]string{"user":"pld","password":"pld"}
			c.JSON(401, gin.H{"status": "unauthorized","message":"you should use this data","data":gin.H{"user":"pld","password":"pld"}})
		}
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func V1IndexLoginfunc(c *gin.Context) {
	name := c.Params.ByName("name")
	//c.Logger().Infof("the param is name=%v",name)
	user:=service.NewUserServiceInstace().GetInfo(name)
	//c.Set("innerName", name)
	//message := getInfo(c)
	//c.String(200, message)
	c.JSON(http.StatusOK,user)
}



func V1IndexSubmitfunc(c *gin.Context) {
	c.String(200, "submit")
}

func V2IndexSubmitfunc(c *gin.Context) {
	user :=make(map[string]string)
	user["name"]="dxc"
	user["address"]="回龙观"
	c.JSON(200,user)
}

func AccessAPIWithAuthorized(c *gin.Context){
	plogger.NewInstance().GetLogger().Info("access the method which need authorized .")

	c.JSON(http.StatusOK, gin.H{
		"secret": "The secret url need to be authorized",
		"status":"success",
	})

}

/**
 * 此处的异常，没有处理，若没有异常中间件处理，数据请求走到此处，会直接异常，没有response；若有了这个异常中间件，可以把异常抛出异常 服务器内部错误。
 */
func  Exception(c *gin.Context){
	panic("something is wrong")


}
