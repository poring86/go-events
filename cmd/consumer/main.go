package main

import (
	"fmt"
	"go-events/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OppenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs)
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
