package utils

import (
	"gin-web-demo/common/dict"
	"strconv"
	"time"
)

/**
 * 获取当前 8 位字符长度的日期
 */
func GetCurrentDate()(dateLen8 string){
	currentDate :=time.Now().Format(dict.SysTimeFmt4compact)[:8]
	return currentDate
}

func GetCurrentTimeUnix() string {
	//当前时间戳
	t1 := time.Now().Unix()
	timeContent := strconv.FormatInt(t1,10)
	return timeContent
}
