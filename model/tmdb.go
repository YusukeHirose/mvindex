package model

type ApiConfig struct {
	Key string `envconfig:"SERVICE_HOST"`
}
