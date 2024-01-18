package configuration

import (
	"github.com/spf13/viper"
)

type config struct {
	Server        serverConfig   `mapstructure:",squash"`
	Database      databaseConfig `mapstructure:",squash"`
	Jwt           jwtConfig      `mapstructure:",squash"`
	Minio         minioConfig    `mapstructure:",squash"`
	RabbitMq      rabbitMqConfig `mapstructure:",squash"`
	Redis         redisConfig    `mapstructure:",squash"`
	Otel          otelConfig     `mapstructure:",squash"`
	Env           string         `mapstructure:"ENV"`
	ServicName    string         `mapstructure:"SERVICE_NAME"`
	GlobalTimeout int            `mapstructure:"GLOBAL_TIMEOUT"`
}

type serverConfig struct {
	Port string `mapstructure:"SERVER.PORT"`
}

type databaseConfig struct {
	Username     string `mapstructure:"DATASOURCE_USERNAME"`
	Password     string `mapstructure:"DATASOURCE_PASSWORD"`
	Host         string `mapstructure:"DATASOURCE_HOST"`
	Port         string `mapstructure:"DATASOURCE_PORT"`
	DbName       string `mapstructure:"DATASOURCE_DB_NAME"`
	PoolMaxConn  int    `mapstructure:"DATASOURCE_POOL_MAX_CONN"`
	PoolIdleConn int    `mapstructure:"DATASOURCE_POOL_IDLE_CONN"`
	PoolLifeTime int    `mapstructure:"DATASOURCE_POOL_LIFE_TIME"`
}

func (d databaseConfig) Dsn() string {
	return "host=" + d.Host + " user=" + d.Username + " password=" + d.Password + " dbname=" + d.DbName + " port=" + d.Port + " sslmode=disable TimeZone=Asia/Jakarta"
}

type jwtConfig struct {
	Secret  string `mapstructure:"JWT_SECRET"`
	Expired int    `mapstructure:"JWT_EXPIRE_MINUTES_COUNT"`
}

type minioConfig struct {
	Endpoint  string `mapstructure:"MINIO_ENDPOINT"`
	Port      string `mapstructure:"MINIO_PORT"`
	AccessKey string `mapstructure:"MINIO_ACCESSKEY"`
	SecretKey string `mapstructure:"MINIO_SECRETKEY"`
	Bucket    string `mapstructure:"MINIO_BUCKET"`
}

type rabbitMqConfig struct {
	Host     string `mapstructure:"RABBITMQ_HOST"`
	Port     string `mapstructure:"RABBITMQ_PORT"`
	Username string `mapstructure:"RABBITMQ_USERNAME"`
	Password string `mapstructure:"RABBITMQ_PASSWORD"`
}

type redisConfig struct {
	Host            string `mapstructure:"REDIS_HOST"`
	Port            string `mapstructure:"REDIS_PORT"`
	Password        string `mapstructure:"REDIS_PASSWORD"`
	PoolMaxSize     int    `mapstructure:"REDIS_POOL_MAX_SIZE"`
	PoolMinIdleSize int    `mapstructure:"REDIS_POOL_MIN_IDLE_SIZE"`
}

type otelConfig struct {
	Endpoint      string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	Insecure      bool   `mapstructure:"OTEL_EXPORTER_OTLP_INSECURE"`
	EnableTracing bool   `mapstructure:"OTEL_ENABLE_TRACING"`
	EnableMetric  bool   `mapstructure:"OTEL_ENABLE_METRIC"`
	EnableLogs    bool   `mapstructure:"OTEL_ENABLE_LOGS"`
}

var viperInstance *viper.Viper
var configInstance config

func Env(filenames ...string) config {
	initViper(filenames...)

	return configInstance
}

func Viper(filenames ...string) *viper.Viper {
	initViper(filenames...)
	return viperInstance
}

func initViper(filenames ...string) {
	if viperInstance == nil {
		viperInstance = viper.New()

		if len(filenames) > 0 {
			viperInstance.SetConfigFile(filenames[0])
		} else {
			viperInstance.SetConfigFile(".env")

		}

		viperInstance.AutomaticEnv()
		viperInstance.ReadInConfig()

		setDefaultViper()

		viperInstance.Unmarshal(&configInstance)

	}
}

func setDefaultViper() {
	// set default value
	viperInstance.SetDefault("SERVICE_NAME", "initial")
	viperInstance.SetDefault("ENV", "development")
	viperInstance.SetDefault("GLOBAL_TIMEOUT", 20)
	viperInstance.SetDefault("SERVER.PORT", "0.0.0.0:8000")
	viperInstance.SetDefault("OTEL_EXPORTER_OTLP_INSECURE", true)
	viperInstance.SetDefault("OTEL_ENABLE_TRACING", true)
	viperInstance.SetDefault("OTEL_ENABLE_METRIC", true)
	viperInstance.SetDefault("OTEL_ENABLE_LOGS", true)
}
