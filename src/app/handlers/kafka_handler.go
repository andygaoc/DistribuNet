package handlers

import (
	"log"

	"app/models"
	"app/services"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaHandler struct {
	MessageService *services.MessageService
}

func (h *KafkaHandler) ConsumeMessages() {
	// 创建Kafka消费者并设置配置
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "your-group-id",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	// 订阅要消费的主题
	err = c.SubscribeTopics([]string{"your-topic"}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	// 消费消息并调用MessageService保存消息
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			h.handleMessage(&msg)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	// 处理错误和关闭消费者
	c.Close()
}

func (h *KafkaHandler) handleMessage(msg *kafka.Message) {
	// 解析消息
	message := &models.Message{
		ID:   int(msg.Headers[0].Value),
		Text: string(msg.Value),
	}

	// 调用MessageService保存消息
	err := h.MessageService.SaveMessage(message)
	if err != nil {
		log.Println("Error saving message:", err)
	}
}
