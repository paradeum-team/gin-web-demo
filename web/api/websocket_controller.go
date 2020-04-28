package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"time"
)

var websocketConn *websocket.Conn

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
// @Tags ws-ping-pong
// @Summary  use websocket
// @Description websocket 用法
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response
// @Router /ws/ping [get]
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	fmt.Printf("time.now=%v \n",time.Now().String())
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)


	if err != nil {
		return
	}
	websocketConn=ws
	//go Job() //启动定时任务，模拟服务端主动向 客户端发送数据。
	go InitiativeWS(websocketConn)//启动异步任务，模拟服务端主动向 客户端发送数据。
	defer ws.Close()
	rand.Seed(time.Now().UnixNano())
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("read message error...")
			break
		}
		fmt.Printf("msg=%s \n",string(message))
		nextSeri:=rand.Intn(1000)
		//if string(message) == "ping" {
			message = []byte(fmt.Sprintf("pong.nextSeri=%v",nextSeri))
		//}

		fmt.Printf("message=%s \n",message)
		//写入ws数据
		err = ws.WriteMessage(mt, message)

		if err != nil {
			fmt.Printf("send.err=%v \n",err)
			break
		}
	}
}