package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var Address = []string{"127.0.0.1:19092"}

func main() {
	syncProducer(Address)
	// aSyncProducer()
}

//同步消息模式
func syncProducer(address []string) {
	// 配置
	config := sarama.NewConfig()
	// 属性设置
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// 创建生成者
	p, err := sarama.NewSyncProducer(address, config)
	// 判断错误
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	// 最后关闭生产者
	defer p.Close()
	// 主题名称
	topic := "topic1"
	// 消息
	srcValue := "sync: this is a message. index=%d"
	// 循环发消息
	for i := 0; i < 10; i++ {
		// 格式化消息
		value := fmt.Sprintf(srcValue, i)
		// 创建消息
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		// 发送消息
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
		}
		// 每隔两秒发送一个消息
		time.Sleep(2 * time.Second)
	}
}

// 异步消息
func aSyncProducer() {

	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer([]string{"192.168.18.128:9092"}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	//循环判断哪个通道发送过来数据.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				//fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message 0606 " + time11.Format("15:04:05")

		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: "topic2",
		}

		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)

		//使用通道发送
		producer.Input() <- msg
	}
}
