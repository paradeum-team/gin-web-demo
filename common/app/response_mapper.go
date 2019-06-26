package app


import (
	"gin-web-demo/common/e"
	"github.com/gin-gonic/gin"

)



type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func NewResponse(C *gin.Context,httpCode, errCode int, data interface{}) {
	C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
