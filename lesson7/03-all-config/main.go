package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type LogLevel string

const (
	defaultLogLevel LogLevel = "error"
	defaultTimeout           = 30 * time.Second
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host     string
	Port     int
	LogLevel LogLevel      `yaml:"log_level"`
	Timeout  time.Duration `yaml:"timeout,omitempty"`
	Options  []string      `yaml:"-"`
}

func main() {
	configFile := flag.String("config", "", "application configuration file")

	flag.Parse()

	if configFile == nil || *configFile == "" {
		log.Fatal("config file not specified")
	}

	data, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("read config file error: %s", err)
	}

	// expand envs
	data = []byte(os.ExpandEnv(string(data)))

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(false)

	cfg := Config{}
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("broken config file: %s", err)
	}

	if cfg.Server.LogLevel == "" {
		cfg.Server.LogLevel = defaultLogLevel
	}

	if cfg.Server.Timeout == 0 {
		cfg.Server.Timeout = defaultTimeout
	}

	fmt.Printf("%+v\n", cfg.Server)
}
