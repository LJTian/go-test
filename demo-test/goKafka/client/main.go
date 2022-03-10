package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

var (
	kafkaConsumer *cluster.Consumer
	kafkaBrokers  = []string{"150.158.78.82:9092"}
	kafkaTopic    = "topic1"
	groupId       = "test_1"
)

func init() {
	// 配置
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	// 创建消费者
	kafkaConsumer, err = cluster.NewConsumer(kafkaBrokers, groupId, []string{kafkaTopic}, config)
	if err != nil {
		panic(err.Error())
	}
	if kafkaConsumer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v, topic: %v, group: %v}\n", kafkaBrokers, kafkaTopic, groupId))
	}
	fmt.Printf("kafka init success, consumer -> %v, topic -> %v, ", kafkaConsumer, kafkaTopic)
}

func main() {
	for {
		select {
		case msg, ok := <-kafkaConsumer.Messages():
			if ok {
				fmt.Printf("kafka 接收到的消息: %s \n", msg.Value)
				kafkaConsumer.MarkOffset(msg, "")
			} else {
				fmt.Printf("kafka 监听服务失败\n")
			}
		case err, ok := <-kafkaConsumer.Errors():
			if ok {
				fmt.Printf("consumer error: %v\n", err)
			}
		case ntf, ok := <-kafkaConsumer.Notifications():
			if ok {
				fmt.Printf("consumer notification: %v\n", ntf)
			}
		}
	}
}
