package main

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello tianlj"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

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

func main() {

	var ramq RabbitMQ
	err := ramq.Connect("10.211.55.3", "5672")
	defer ramq.CloseConnect()

	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return
	}

	err = ramq.ExchangeCreat("fanout", "TranMsg")
	if err != nil {
		failOnError(err, "Failed exchange")
	}

	go func() {
		ch := ramq.channelMap["TranMsg"]
		q, err := ramq.channelMap["TranMsg"].QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = ch.QueueBind(
			q.Name,    // queue name
			"",        // routing key
			"TranMsg", // exchange
			false,
			nil,
		)
		failOnError(err, "Failed to bind a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf(" [%s]worker %s", "TranMsg", d.Body)
			}
		}()

		log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
		<-forever
	}()

	forever1 := make(chan bool)

	for i := 0; i < 100; i++ {
		body := fmt.Sprintf("hello %d", i)
		err = ramq.Send("TranMsg", []byte(body))
		if err != nil {
			failOnError(err, "Failed  Send")
		}
		time.Sleep(1 * time.Second)
	}

	<-forever1
	return
}
