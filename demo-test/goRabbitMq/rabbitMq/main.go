package main

import (
	"fmt"
	"log"
	"time"
)

func worker(ramq *RabbitMQ, topic string, routingKey string, name string) {

	q, err := ramq.CreatQueue(topic, "")
	if err != nil {
		failOnError(err, "Failed to declare a queue")
		return
	}

	err = ramq.BindQueue(topic, q.Name, routingKey)
	if err != nil {
		failOnError(err, "Failed to bind a queue")
		return
	}

	msgs, err := ramq.CreatConsume(topic, q.Name, "", true)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("name:[%s] ---- [%s-%s]worker %s", name, topic, routingKey, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func producer(ramq *RabbitMQ, topic string, name string) {

	var err error

	// 发送数据
	forever1 := make(chan bool)

	for i := 0; i < 100; i++ {
		body := fmt.Sprintf("hello [%d]", i)
		if i%2 == 0 {
			err = ramq.Send(topic, "MngTran.Log", []byte(body))
		} else {
			err = ramq.Send(topic, "AcctTran.Log", []byte(body))
		}
		log.Printf("name[%s]  -- Say[%s]", name, body)

		if err != nil {
			failOnError(err, "Failed  Send")
		}
		time.Sleep(1 * time.Second)
	}

	<-forever1
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

	err = ramq.CreatExchange("topic", topic)
	if err != nil {
		failOnError(err, "Failed exchange")
	}

	// 开启消费者
	go worker(&ramq, topic, "MngTran.*", "worker1")
	go worker(&ramq, topic, "*.Log", "worker2")
	go worker(&ramq, topic, "*.Log", "worker3")
	go worker(&ramq, topic, "AcctTran.*", "worker4")

	// 开启生产者
	go producer(&ramq, topic, "producer1")
	go producer(&ramq, topic, "producer2")

	stopCh := make(chan int)

	<-stopCh
	return
}
