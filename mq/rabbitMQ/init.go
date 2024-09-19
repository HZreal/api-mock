package rabbitMQ

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/config"
	"gin-init/constant"
	"gin-init/consumers"
	"github.com/streadway/amqp"
	"log"
)

var Conn *amqp.Connection

var producer *RabbitMQProducer

var consumer *RabbitMQConsumer

/*
* 生产者
 */
func createGlobalProducer() {
	producer = NewProducer(Conn)
}

/*
* 消费者
 */

// 注册器
var registers = []struct {
	topic   constant.Topic
	handler func(msg []byte)
}{
	{topic: constant.DemoTopic, handler: consumers.HandleMessage1},
	{topic: constant.FirstTopic, handler: consumers.HandleMessage2},
}

func startConsumers() {
	consumer = NewConsumer(Conn)

	for _, register := range registers {
		// 启动消费者
		err := consumer.Listen(register.topic, register.handler)
		if err != nil {
			log.Fatalf("Failed to start consumer for %v: %v", register.topic, err)
		}
	}

}

func Start() {
	var err error
	Conn, err = amqp.Dial(config.Conf.RabbitMQ.GetUrl())
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ")

	//
	createGlobalProducer()

	//
	startConsumers()
}
