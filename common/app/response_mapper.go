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
type BadResponse struct {
	Code int         `json:"code"`
	Error  string      `json:"error"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func NewResponse(C *gin.Context,httpCode, errCode int, data interface{}) {
	C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
// Response setting gin.JSON
func NewBadResponse(C *gin.Context,httpStatusCode, errCode int, data interface{}) {
	C.JSON(httpStatusCode, BadResponse{
		Code: errCode,
		Error:  e.GetMsg(errCode),
		Data: data,
	})
	return
}