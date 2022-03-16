package main

import (
	"errors"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn       *amqp.Connection
	ip         string
	port       string
	channelMap map[string]*amqp.Channel
}

// 建立链接
func (ramq *RabbitMQ) Connect(ip string, port string) (err error) {

	ramq.ip = ip
	ramq.port = port
	url := "amqp://guest:guest@" + ramq.ip + ":" + ramq.port + "/"
	ramq.conn, err = amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	return
}

// 关闭链接
func (ramq *RabbitMQ) CloseConnect() (err error) {
	ramq.CloseChannels()
	return ramq.conn.Close()
}

// 创建转换器
func (ramq *RabbitMQ) CreatExchange(exchangeType string, topic string) (err error) {

	if ramq.channelMap == nil {
		ramq.channelMap = make(map[string]*amqp.Channel)
	} else {
		if _, ok := ramq.channelMap[topic]; ok {
			return errors.New(topic + "已存在")
		}
	}

	ch, err := ramq.conn.Channel()
	if err != nil {
		return
	}
	err = ch.ExchangeDeclare(
		topic,        // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	ramq.channelMap[topic] = ch
	return
}

// 关闭管道
func (ramq *RabbitMQ) CloseChannels() {
	for k, _ := range ramq.channelMap {
		ramq.CloseChannelByTopic(k)
	}
}

// 创建队列
func (ramq *RabbitMQ) CreatQueue(topic string, QueueName string) (queue *amqp.Queue, err error) {

	q, err := ramq.channelMap[topic].QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	queue = &q
	return
}

// 绑定队列
func (ramq *RabbitMQ) BindQueue(topic string, queueName string, routingKey string) (err error) {
	err = ramq.channelMap[topic].QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		topic,      // exchange
		false,
		nil,
	)

	return
}

// 基于事件主题 关闭管道
func (ramq *RabbitMQ) CloseChannelByTopic(topic string) (err error) {
	if v, ok := ramq.channelMap[topic]; ok {
		v.ExchangeDelete(topic, false, true)
		v.Close()
		return
	}
	return errors.New(topic + "不存在！")
}

// 发送消息
func (ramq *RabbitMQ) Send(topic string, routingKey string, msg []byte) (err error) {

	if _, ok := ramq.channelMap[topic]; !ok {
		return errors.New(topic + "不存在！")
	}

	ch := ramq.channelMap[topic]
	err = ch.Publish(
		topic,      // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	return
}

//
func (ramq *RabbitMQ) CreatConsume(topic string, queueName string, ConsumeName string, aotoAck bool) (msgs <-chan amqp.Delivery, err error) {

	msgs, err = ramq.channelMap[topic].Consume(
		queueName,   // queue
		ConsumeName, // consumer
		aotoAck,     // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	return
}
