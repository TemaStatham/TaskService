package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	_ "github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"os"
)

// Наша структура для сообщения
type Message struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  int16  `json:"role"`
}

func main() {
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
		os.Exit(1)
	}

	// Dummy Data
	userId := [5]int{100001, 100002, 100003, 100004, 100005}
	emails := [5]string{"POST00001", "POST00002", "POST00003", "POST00004", "POST00005"}
	roles := [5]int16{1, 2, 3, 4, 5}

	for {
		// we are going to take random data from the dummy data
		message := Message{
			Id:    userId[rand.Intn(len(userId))],
			Email: emails[rand.Intn(len(emails))],
			Role:  roles[rand.Intn(len(roles))],
		}

		jsonMessage, err := json.Marshal(message)

		if err != nil {
			log.Fatalln("Failed to marshal message:", err)
			os.Exit(1)
		}

		msg := &sarama.ProducerMessage{
			Topic: "users_topic",
			Value: sarama.StringEncoder(jsonMessage),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("Failed to send message:", err)
			os.Exit(1)
		}
		log.Println("Message sent!")
	}
}
