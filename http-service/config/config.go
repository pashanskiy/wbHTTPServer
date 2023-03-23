package config

type Config struct {
	HTTPServerHost string `env:"HTTP_SERVER_HOST" envDefault:"0.0.0.0"`
	HTTPServerPort int    `env:"HTTP_SERVER_PORT" envDefault:"9001"`

	LogRootLevel string `env:"LOG_ROOT_LEVEL" envDefault:"TRACE"`
}
