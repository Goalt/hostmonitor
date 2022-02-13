package config

import "time"

type Config struct {
	Logger      Logger
	GRPCServer  GRPCServer
	ProxyServer ProxyServer
	Updater     Updater
	Host        Host
}

type GRPCServer struct {
	Port int
}

type ProxyServer struct {
	Port int
}

type Logger struct {
	SetReportCaller bool
	Level           int
}

type Updater struct {
	Interval time.Duration
}

type Host struct {
	User     string
	Password string
	Address  string
}
