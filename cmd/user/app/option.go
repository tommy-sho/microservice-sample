package app

type Option struct {
	Port                 string `env:"GRPC_PORT" envDefault:"50051"`
	ConnectTimeout       uint   `env:"GRPC_CONNECT_TIMEOUT" envDefault:"1000"`
	KeepAlive            uint   `env:"GRPC_KEEP_ALIVE_MS" envDefault:"5000"` // 5000ms = 5sec
	GracefulPeriod       uint   `env:"GRACEFUL_PERIOD" envDefault:"10000"`   // 10000 = 10sec
	MetricsAdnHealthPort string `env:"METRICS_HEALTH_PORT" envDefault:"8080"`
	LogLevel             string `env:"LOG_LEVEL" envDefault:"info"`
}
