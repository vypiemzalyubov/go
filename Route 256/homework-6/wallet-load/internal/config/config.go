package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/keepalive"
)

var (
	debugHost        = "localhost"
	dockerHost       = "0.0.0.0"
	grpcPort         = 8002
	gatewayPort      = 8001
	metricsPort      = 9100
	statusPort       = 8000
	DSNKey           = "PG_DSN"             //nolint
	ConsumeTickerKey = "KAFKA_CONSUME_TICK" //nolint
	CbrURL           = "CBR_URL"            //nolint

	debug = flag.Bool("debug", false, "Application host: localhost")
)

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Addr string
	Path string
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Addr             string
	WithReflection   bool
	ServerParameters keepalive.ServerParameters
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Addr string
}

// Database - contains all parameters database connection.
type Database struct {
	Dsn               string
	Driver            string
	ConnTimeout       time.Duration
	MaxOpenConn       int
	MasterMaxIdleConn int
	Migrations        string
}

// Kafka - contains all parameters kafka connection.
type Kafka struct {
	Topic         string
	Hosts         []string
	ConsumeTicker time.Duration
}

// Status config for service.
type Status struct {
	Addr          string
	LivenessPath  string
	ReadinessPath string
}

// AppConfig app config
type AppConfig struct {
	Debug    bool
	Metrics  Metrics
	Grpc     Grpc
	Rest     Rest
	Database Database
	Kafka    Kafka
	Status   Status
	Jobs     JobsSetting
}

type JobsSetting struct {
	DurationCollectOperation time.Duration
}

// InitAppConfig load app config
func InitAppConfig() *AppConfig {
	flag.Parse()

	host := dockerHost
	if *debug {
		host = debugHost
	}

	return &AppConfig{
		Debug: *debug,
		Database: Database{
			Dsn:               os.Getenv(DSNKey),
			ConnTimeout:       time.Duration(10),
			MaxOpenConn:       50,
			MasterMaxIdleConn: 50,
			Driver:            "pgx",
			Migrations:        "migrations",
		},
		Kafka: Kafka{
			Topic:         "wallet_external_operations",
			Hosts:         []string{"kafka:9092"},
			ConsumeTicker: parseDuration(os.Getenv(ConsumeTickerKey)),
		},
		Grpc: Grpc{
			Addr: fmt.Sprintf("%s:%v", host, grpcPort),
			ServerParameters: keepalive.ServerParameters{
				MaxConnectionIdle: time.Duration(1) * time.Minute,
				Timeout:           time.Duration(30) * time.Second,
				MaxConnectionAge:  time.Duration(1) * time.Minute,
				Time:              time.Duration(1) * time.Minute,
			},
		},
		Rest: Rest{
			Addr: fmt.Sprintf("%s:%v", host, gatewayPort),
		},
		Metrics: Metrics{
			Addr: fmt.Sprintf("%s:%v", host, metricsPort),
			Path: "/metrics",
		},
		Status: Status{
			Addr:          fmt.Sprintf("%s:%v", host, statusPort),
			LivenessPath:  "/live",
			ReadinessPath: "/ready",
		},
		Jobs: JobsSetting{
			DurationCollectOperation: time.Duration(30) * time.Second,
		},
	}
}

func parseDuration(dur string) time.Duration {
	d, err := time.ParseDuration(dur)
	if err != nil {
		return 30 * time.Second
	}

	return d
}
