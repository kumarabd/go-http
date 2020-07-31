package logger

import (
	"github.com/kumarabd/gokit/logger"
)

func New(appname string) logger.LogHandler {
	l, err := logger.New(appname)
	if err != nil {
		return nil
	}
	return l
}
