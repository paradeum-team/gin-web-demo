package plogger

import (
	"gin-web-demo/common/dict"
	"gin-web-demo/common/utils"
	"fmt"
	pldconf "gin-web-demo/config"
	"github.com/kataras/golog"
	"log"
	"os"
	"path/filepath"
)

type PldLogger struct {
	logger      *golog.Logger
	currentDate string //当前时间
}

var pldLoggerInstance *PldLogger

func NewInstance() *PldLogger {
	currentDate := utils.GetCurrentDate() //当前的8位长度的日期
	pldLoggerInstance = &PldLogger{
		logger:      golog.Default,
		currentDate: currentDate,
	}

	baseLogPath := filepath.Join(pldconf.AppConfig.Server.DataPath, dict.LOG_FOLDER)
	//check base log dir .if not exits then create .
	createAfsLogDir(baseLogPath)

	logFileName := "sys_" + currentDate + ".log"
	logFilePath := filepath.Join(baseLogPath, logFileName)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logFilePath, err))
		return nil
	}

	//设置output
	pldLoggerInstance.logger.SetOutput(os.Stdout)
	pldLoggerInstance.logger.AddOutput(f)
	return pldLoggerInstance
}

func (lf *PldLogger) GetLogger() *golog.Logger {
	if pldLoggerInstance == nil {
		NewInstance()
	} else {
		if lf.currentDate == utils.GetCurrentDate() {
			//同一天，说明日志不用切换文件，否则就新打开一个文件
		} else {
			NewInstance()
		}
	}
	return lf.logger
}

func createAfsLogDir(baseLogPath string) {

	err := os.MkdirAll(baseLogPath, os.ModePerm) //创建多级目录，如果path已经是一个目录，MkdirAll什么也不做，并返回nil。

	if err != nil {
		log.Println("ERROR: init log dir is something wrong ...")
		os.Exit(1) //日志文件目录创建不成功，则失败
	}

}

/**
var loggerFactoryInstance *loggerFactory

type loggerFactory struct {
	logger *golog.Logger
}

func NewInstance() *loggerFactory {
	loggerFactoryInstance=&loggerFactory{
		logger:golog.Default,
	}
	baseLogPath :=filepath.Join(pldconf.AppConfig.Server.APIDataFolder,dict.RN_LOG_FOLDER)
	//check base log dir .if not exits then create .
	createAfsLogDir(baseLogPath)

	logFileName :="rn_" + time.Now().Format("20060102150405")[:8] + ".log"
	logFilePath := filepath.Join(baseLogPath,logFileName)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

		log.Printf("ERROR: %s\n", fmt.Sprintf("%s append|create failed:%v", logFilePath, err))
		return nil
	}

	loggerFactoryInstance.GetLogger().SetOutput(os.Stdout)
	loggerFactoryInstance.GetLogger().AddOutput(f)
	return loggerFactoryInstance
}
**/
