package main

import (
	"encoding/json"
)

// This is just the init and glue code. Connects up the separate pieces.

func kafkaGlue(fromKafka chan string, toClients chan []byte) {
	for x := range fromKafka {
		dat, _ := lookup(parse(x))
		js, _ := json.Marshal(dat)
		toClients <- js
	}
}

func main() {
	// get the configuration from config.yaml
	cfg := GetConfig()
	// start up websockets
	hub := NewHub()
	// build the Kafka channels
	fromKafka := make(chan string)
	toClients := hub.broadcast
	// start up the goroutines to take care of kafka and websockets
	go startKafkaConnection(cfg.Kafka.Hosts, 0, cfg.Kafka.Topic, fromKafka)
	go hub.run()
	go kafkaGlue(fromKafka, toClients)
	// start serving
	StartWeb(cfg.Http.Port, hub)
}
