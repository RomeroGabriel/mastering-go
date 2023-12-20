package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int32
	Msg string
}

var countId int32 = 0

func rabbitMQMessage(ch chan<- Message) {
	// Infity loop
	for {
		atomic.AddInt32(&countId, 1)
		msg := Message{countId, fmt.Sprintf("Message %d from RabbitMQ", countId)}
		ch <- msg
		if countId > 8 {
			time.Sleep(4 * time.Second)
		}
	}
}

func kafkaMessage(ch chan<- Message) {
	// Infity loop
	for {
		atomic.AddInt32(&countId, 1)
		msg := Message{countId, fmt.Sprintf("Message %d from Kafka", countId)}
		ch <- msg
		if countId > 4 {
			time.Sleep(4 * time.Second)
		}
	}
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	go rabbitMQMessage(ch1)
	go kafkaMessage(ch2)

	for i := 0; i < 20; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from ch1: ID: %d - %s\n", msg.id, msg.Msg)
		case msg := <-ch2:
			fmt.Printf("Received from ch2: ID: %d - %s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 2):
			log.Print("timeout")
		}
	}

}
