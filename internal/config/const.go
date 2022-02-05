package config

import (
	gormLog "gorm.io/gorm/logger"
)

const (
	InfoLevel  = 1
	WarnLevel  = 2
	ErrorLevel = 3

	GormLogLevel = gormLog.Error
)