package apollo

import "github.com/micro/micro/v3/service/logger"

type DefaultLogger struct {
}

func (DefaultLogger) Debugf(format string, params ...interface{}) {
	logger.Debugf(format, params...)
}

func (DefaultLogger) Infof(format string, params ...interface{}) {
	logger.Infof(format, params...)
}

func (DefaultLogger) Warnf(format string, params ...interface{}) {
	logger.Warnf(format, params...)
}

func (DefaultLogger) Errorf(format string, params ...interface{}) {
	logger.Errorf(format, params...)
}

func (DefaultLogger) Debug(v ...interface{}) {
	logger.Debug(v...)
}
func (DefaultLogger) Info(v ...interface{}) {
	logger.Info(v...)
}

func (DefaultLogger) Warn(v ...interface{}) {
	logger.Warn(v...)
}

func (DefaultLogger) Error(v ...interface{}) {
	logger.Error(v...)
}
