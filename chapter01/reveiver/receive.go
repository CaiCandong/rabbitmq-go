package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 建立TCP Scoket连接
	//Scoket连接对象是Rabbit中的一个v-host
	// v - host类似连接到Mysql的某个数据库
	conn, err := amqp.Dial("amqp://test:test@localhost:5672/center")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 在TCP连接上建立一个信道channel
	// 一个TCP连接上可以建立多条信号,且彼此独立
	// 能够起到复用TCP连接的作用
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 在信道上建立消息队列
	// 这里没有使用交换机
	// 发送方和接受方的消息队列要保持一致
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack 自动消息确认
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
