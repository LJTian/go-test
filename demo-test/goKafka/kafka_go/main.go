package main

import (
	kafka "github.com/segmentio/kafka-go"
)

func main() {

	// to produce messages

	//var controllerConn *kafka.Conn
	//controllerConn, err := kafka.Dial("tcp", "10.211.55.3:9092")
	//
	//controllerConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//_, err = controllerConn.WriteMessages(
	//	kafka.Message{Topic: "my-topic", Value: []byte("one!")},
	//	kafka.Message{Topic: "my-topic", Value: []byte("two!")},
	//	kafka.Message{Topic: "my-topic", Value: []byte("three!")},
	//)
	//if err != nil {
	//	log.Fatal("failed to write messages:", err)
	//}
	//
	//if err := controllerConn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}
	CreateTopics()
}

// 创建主题
func CreateTopics() {
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "127.0.0.1:19092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", "127.0.0.1:19092")
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

}
