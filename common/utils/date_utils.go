package utils

import (
	"gin-web-demo/common/dict"
	"time"
)

/**
 * 获取当前 8 位字符长度的日期
 */
func GetCurrentDate()(dateLen8 string){
	currentDate :=time.Now().Format(dict.SysTimeFmt4compact)[:8]
	return currentDate
}