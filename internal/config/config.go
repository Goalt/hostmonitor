package config

type Config struct {
	Logger     Logger
	GRPCServer GRPCServer
	ProxyServer ProxyServer
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