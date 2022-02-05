package config

import "fmt"

type Config struct {
	Logger      Logger
}

type Logger struct {
	SetReportCaller bool
	Level           int
}