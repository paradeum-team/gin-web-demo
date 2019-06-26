package controller

import (
	"gin-web-demo/common/app"
	"gin-web-demo/common/e"
	"gin-web-demo/common/utils"
	"gin-web-demo/service"
	"gin-web-demo/vo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary user login with username and pwd
// @Description user login  with username and pwd
// @Accept  json
// @Produce  json
// @Param user body vo.LoginJSON true "user body"
// @Success 200 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/login [post]
func Login(c *gin.Context) {
	plogger.NewInstance().GetLogger().Info("user  login, will access ...")

	var json vo.LoginJSON
	// If EnsureBody returns false, it will write automatically the error
	// in the HTTP stream and return a 400 error. If you want custom error
	// handling you should use: c.ParseBody(interface{}) error
	log.Println("come here 123")

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User == "pld" && json.Password == "pld" {
			app.NewResponse(c, http.StatusOK, e.SUCCESS, gin.H{"status": "you are logged in", "ext": "ext345"})
		} else {
			//data :=map[string]string{"user":"pld","password":"pld"}
			app.NewResponse(c, http.StatusUnauthorized, e.ERR_NO_AUTH, gin.H{"status": "unauthorized", "message": "you should use this data", "data": gin.H{"user": "pld", "password": "pld"}})
		}
	} else {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		app.NewResponse(c, http.StatusInternalServerError, e.ERROR, gin.H{"error": err.Error()})
		return
	}
}

// User   godoc
// @Summary find user by name
// @Description find user  by name
// @Accept  json
// @Produce  json
// @Param name path string true "user name"
// @Success 200 {object} entity.User
// @Router /v1/users/{name} [get]
func GetUserByName(c *gin.Context) {
	name := c.Params.ByName("name")
	//c.Logger().Infof("the param is name=%v",name)
	user := service.NewUserServiceInstace().GetInfo(name)
	//c.Set("innerName", name)
	//message := getInfo(c)
	//c.String(200, message)
	c.JSON(http.StatusOK,user)
	//app.NewResponse(c, http.StatusOK, e.SUCCESS, user)
}


// @Summary list all users
// @Description list all users
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response "[]entity.User"
// @Router /v1/users [get]
func ListUsers(c *gin.Context) {
	userList :=service.NewUserServiceInstace().ListUsers()
	app.NewResponse(c,http.StatusOK,e.SUCCESS,userList)
}

// User login godoc
// @Summary  router 的另一种写法
// @Description 测试数据：router 的另一中写法
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.User
// @Router /v2/submit [get]
func GetTestUserData(c *gin.Context) {
	user := make(map[string]string)
	user["name"] = "dxc"
	user["address"] = "回龙观"
	c.JSON(200, user)
}



// User login godoc
// @Summary  Basic auth 使用方法
// @Description 在header 头增加 ["Authorization":"Basic Zm9vOmJhcg=="],["Authorization":"Basic YWRtaW46cGFzc3dvcmQ="]
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response
// @Router /auth/secret [get]
func AccessAPIWithAuthorized(c *gin.Context) {
	plogger.NewInstance().GetLogger().Info("access the method which need authorized .")

	c.JSON(http.StatusOK, gin.H{
		"secret": "The secret url need to be authorized",
		"status": "success",
	})

}

/**
 * 此处的异常，没有处理，若没有异常中间件处理，数据请求走到此处，会直接异常，没有response；若有了这个异常中间件，可以把异常抛出异常 服务器内部错误。
 */
func Exception(c *gin.Context) {
	panic("something is wrong")

}
