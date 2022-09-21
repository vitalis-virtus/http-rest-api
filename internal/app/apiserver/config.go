package apiserver

import "github.com/vitalis-virtus/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BinAddr  string `toml:"bind_addr"` // the address where we run our web server
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BinAddr:  "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
