package rabbitMQ

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: producer.go
 * @Description:
 */

import (
	"fmt"
	"gin-init/config"
	"gin-init/constant"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type Producer interface {
	Publish(queueName string, body string) error
}

type RabbitMQProducer struct {
	Conn *amqp.Connection
}

func NewProducer(conn *amqp.Connection) *RabbitMQProducer {
	return &RabbitMQProducer{Conn: conn}
}

func (p *RabbitMQProducer) Publish(queueName string, body string) error {
	ch, err := p.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

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

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf("Message sent to queue %s: %s", queueName, body)
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////////

type RabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	once    sync.Once
	url     string
}

// NewRabbitMQService 初始化 RabbitMQService
func NewRabbitMQService() *RabbitMQService {
	url := config.Conf.RabbitMQ.GetUrl()
	return &RabbitMQService{
		url: url,
	}
}

// connect 建立与 RabbitMQ 的连接
func (r *RabbitMQService) connect() error {
	var err error
	r.once.Do(func() {
		r.conn, err = amqp.Dial(r.url)
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}
		r.channel, err = r.conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
	})
	return err
}

// reconnect 用于重连 RabbitMQ
func (r *RabbitMQService) reconnect() {
	r.once = sync.Once{} // 重新设置 once 以便重新连接
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	if err := r.connect(); err != nil {
		log.Fatalf("Failed to reconnect to RabbitMQ: %v", err)
	}
}

func (r *RabbitMQService) Publish(exchange, routingKey string, body []byte) error {
	// 确保连接已建立
	if r.conn == nil || r.channel == nil {
		// TODO 获取连接对象去全局单例里获取
		if err := r.connect(); err != nil {
			return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
		}
	}

	// 声明交换机，指定 direct 模式
	err := r.channel.ExchangeDeclare(
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

	// 发布消息
	err = r.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent, // 持久化消息
			Priority:     0,               // 消息优先级
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		r.reconnect() // 连接失败时尝试重连
		return err
	}

	log.Printf("Message sent: %s", body)
	return nil
}

// ////////////////////////////////////// 以下是业务相关的发送函数 //////////////////////////////////////////////
func (r *RabbitMQService) SendMQMsq111(body []byte) {
	topic := constant.DemoTopic
	_ = r.Publish(topic.Exchange, topic.Route, body)
}

func (r *RabbitMQService) SendMQMsg222(body []byte) {
	topic := constant.FirstTopic
	_ = r.Publish(topic.Exchange, topic.Route, body)
}
