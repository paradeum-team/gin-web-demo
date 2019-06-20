package plogger

import (
	"github.com/kataras/golog"
)

var loggerFactoryInstance *loggerFactory

type loggerFactory struct {
	logger *golog.Logger
}

func NewInstance() *loggerFactory {
	loggerFactoryInstance=&loggerFactory{
		logger:golog.Default,
	}
	return loggerFactoryInstance
}

func (lf *loggerFactory) GetLogger() *golog.Logger {
	if loggerFactoryInstance ==nil{
		NewInstance()
	}
	return loggerFactoryInstance.logger
}
