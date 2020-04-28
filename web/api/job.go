package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"time"
)

var idx int = 0

func Job() {
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	task := func() {
		fmt.Println("hello world", time.Now())
		if websocketConn != nil {
			idx += 1
			websocketConn.WriteMessage(websocket.TextMessage, []byte(time.Now().String()))
			fmt.Printf("idx=%d \n", idx)
			if idx > 10 {
				websocketConn.Close()
				websocketConn = nil
			}
		}
	}
	//定时任务
	spec := "*/5 * * * * ?" //cron表达式，每五秒一次
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止
}

/*
 * 模拟websocket 主动把数据传给客户端
 */
func InitiativeWS(wsConn *websocket.Conn) {
	for idx := 0; idx < 10; idx++ {
		if wsConn == nil {
			break
		}
		time.Sleep(time.Second * 1)
		content := fmt.Sprintf("ws.server.send[%d]=%v", idx, time.Now().String())
		fmt.Printf("%s\n", content)
		wsConn.WriteMessage(websocket.TextMessage, []byte(content))
	}
	if wsConn != nil {
		fmt.Printf("clost ws \n")
		wsConn.Close()
		wsConn = nil
	}
}
