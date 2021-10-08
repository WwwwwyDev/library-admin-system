package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	Postgresql struct {
		User     string
		Password string
		Host     string
		Port     int
		Name     string
	}
}
