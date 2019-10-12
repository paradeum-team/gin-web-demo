package utils

import "time"

/**
 * 获取当前 8 位字符长度的日期
 */
func GetCurrentDate()(dateLen8 string){
	currentDate :=time.Now().Format("20060102150405")[:8]
	return currentDate
}