package config

import (
	"encoding/base64"
	"log/slog"
)

type Config struct {
	APIBase    string             `envconfig:"api_base" required:"true"`
	AccessKey  string             `envconfig:"access_key" split_case:"true" required:"true"`
	SecretKey  Base64EncodedValue `envconfig:"secret_key" required:"true"`
	Backend    string             `default:"tesseract"`
	LogLevel   slog.Level         `envconfig:"log_level" default:"INFO"`
	WorkerPort int                `default:"8080"`
}

type Base64EncodedValue []byte

func (b *Base64EncodedValue) Decode(value string) error {
	v, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}

	*b = v

	return nil
}
