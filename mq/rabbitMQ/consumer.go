package rabbitMQ

import (
	"fmt"
	"gin-init/constant"
	"github.com/streadway/amqp"
	"log"
)

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer.go
 * @Description:
 */

type RabbitMQConsumer struct {
	Conn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		Conn: conn,
	}
}

func (c *RabbitMQConsumer) Listen(topic constant.Topic, handler func(body []byte)) error {
	queueName, exchange, routingKey := topic.Queue, topic.Exchange, topic.Route

	ch, err := c.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return err
	}
	// defer ch.Close()

	// 声明交换机，指定 direct 模式
	err = ch.ExchangeDeclare(
		exchange, // exchange name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return err
	}

	_ = ch.QueueBind(queueName, routingKey, exchange, false, nil)

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v", err)
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf("Listening for messages on queue: %s", queueName)
	return nil
}
