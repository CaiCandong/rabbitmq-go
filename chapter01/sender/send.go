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
	//Scoket连接对象是Rabbit中的一个v-host,类似连接到Mysql的某个数据库
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
	// 声明队列是幂等的——仅当队列不存在时才创建
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		// 默认交换机（default exchange)上
		// 是一个由消息代理预先声明好的没有名字（名字为空字符串）
		// 的直连交换机（direct exchange）。
		"",     // exchange 名字为空字符串的直连交换机
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
