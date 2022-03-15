package main

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
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
func (ramq *RabbitMQ) ExchangeCreat(exchangeType string, topic string) (err error) {

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

// 基于事件主题 关闭管道
func (ramq *RabbitMQ) CloseChannelByTopic(topic string) (err error) {
	if v, ok := ramq.channelMap[topic]; ok {
		v.Close()
		return
	}
	return errors.New(topic + "不存在！")
}

// 发送消息
func (ramq *RabbitMQ) Send(topic string, msg []byte) (err error) {

	if _, ok := ramq.channelMap[topic]; !ok {
		return errors.New(topic + "不存在！")
	}

	ch := ramq.channelMap[topic]
	err = ch.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	log.Printf(" [%s] Sent %s", topic, string(msg))
	return
}
