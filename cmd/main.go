package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TemaStatham/TaskService/internal/app"
	"github.com/TemaStatham/TaskService/internal/config"
	"log"
	"math/rand"
	"os"
	"time"
)

// go run ./cmd/main.go --config="./config.yaml"

func main() {
	cfg := config.MustLoad()

	a := app.New()

	//go kafka(*cfg)

	a.MustRun(cfg)
}

type Message struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  int16  `json:"role"`
}

func kafka(cfg config.Config) {
	brokers := []string{cfg.KConfig.Host + ":" + cfg.KConfig.Port}

	configSarama := sarama.NewConfig()
	configSarama.Producer.Return.Successes = true
	configSarama.Version = sarama.V2_0_0_0
	configSarama.Consumer.Offsets.AutoCommit.Enable = true
	configSarama.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	producer, err := sarama.NewSyncProducer(brokers, configSarama)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
		os.Exit(1)
	}
	fmt.Println("Config kafka")

	// Dummy Data
	userId := [5]int{100001, 100002, 100003, 100004, 100005}
	emails := [5]string{"POST00001", "POST00002", "POST00003", "POST00004", "POST00005"}
	roles := [5]int16{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		fmt.Println("Waiting for messages")
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
