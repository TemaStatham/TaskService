package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TemaStatham/TaskService/internal/config"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/internal/repository/postgres"
	"log"
	"time"
)

type consumerGroupHandler struct {
	postgres.UserRepository
}

func newConsumer(userRepository postgres.UserRepository) consumerGroupHandler {
	return consumerGroupHandler{
		userRepository,
	}
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var user model.UserModel
		err := json.Unmarshal([]byte(msg.Value), &user)
		if err != nil {
			fmt.Println("Ошибка парсинга JSON: %v", err)
		}

		err = h.UserRepository.Create(context.Background(), &user)
		if err != nil {
			fmt.Println("Ошибка сохранения: %v", err)
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

type KafkaService struct {
	postgres.UserRepository
}

func NewKafkaService(repository postgres.UserRepository) *KafkaService {
	return &KafkaService{
		repository,
	}
}

func (k *KafkaService) StartConsume(cfg config.KafkaConfig) {
	brokers := []string{cfg.Host + ":" + cfg.Port}
	groupID := cfg.GroupsId

	configSarama := sarama.NewConfig()
	configSarama.Version = sarama.V2_0_0_0
	configSarama.Consumer.Offsets.AutoCommit.Enable = true
	configSarama.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, configSarama)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	go func() {
		for {
			err := consumerGroup.Consume(context.Background(), []string{cfg.UsersTopic}, newConsumer(k.UserRepository))
			if err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
		}
	}()
}
