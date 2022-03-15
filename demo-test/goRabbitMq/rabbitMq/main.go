package main

import (
	"fmt"
	"log"
	"time"
)

func worker(ramq *RabbitMQ, topic string) {
	ch := ramq.channelMap[topic]
	q, err := ramq.channelMap[topic].QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		topic,  // exchange
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
			log.Printf(" [%s]worker %s", topic, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func main() {

	topic := "TranMsg"

	var ramq RabbitMQ
	err := ramq.Connect("10.211.55.3", "5672")
	defer ramq.CloseConnect()

	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return
	}

	err = ramq.ExchangeCreat("fanout", topic)
	if err != nil {
		failOnError(err, "Failed exchange")
	}

	// 开启消费者
	go worker(&ramq, topic)

	// 发送数据
	forever1 := make(chan bool)

	for i := 0; i < 100; i++ {
		body := fmt.Sprintf("hello %d", i)
		err = ramq.Send(topic, []byte(body))
		if err != nil {
			failOnError(err, "Failed  Send")
		}
		time.Sleep(1 * time.Second)
	}

	<-forever1
	return
}
