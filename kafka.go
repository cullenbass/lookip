package main

import (
	"gopkg.in/Shopify/sarama.v1"
	"log"
)

var (
	consumer     sarama.Consumer
	partConsumer sarama.PartitionConsumer
	messages     chan *sarama.ConsumerMessage
)

func startKafkaConnection(brokers []string, partition int32, topic string, out chan string) {
	// configure the Kafka client
	config := sarama.NewConfig()
	config.ClientID = ("lookip")

	// start up the connections
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal("Kafka connection error: ", err)
	}
	defer consumer.Close()
	partConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Kafka topic error: ", err)
	}
	defer partConsumer.Close()

	// start consuming!
	for msg := range partConsumer.Messages() {
		out <- string(msg.Value)
	}
	close(out)
}
