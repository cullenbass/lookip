package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Split these out into multiple structs because reasons. I dunno.
// Not sure I like how it looks over the massive struct.
type Config struct {
	Http  HttpConfig
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Hosts []string `yaml:",flow"`
	Topic string
}

type HttpConfig struct {
	Port int
}

// Should run once at the beginning of runtime, but could theoretically
// run at any point to get an update of the configs.
func GetConfig() Config {
	//configuration
	var cfgFile []byte
	var cfg Config
	cfg = Config{
		Kafka: KafkaConfig{
			Hosts: []string{"127.0.0.1:9092"},
			Topic: "mapping",
		},
		Http: HttpConfig{
			Port: 8085,
		},
	}
	cfgFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("Error parsing configuration file, using defaults.")
	}
	yaml.Unmarshal(cfgFile, &cfg)
	return cfg
}
