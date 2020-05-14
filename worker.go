package main

import (
	"log"
	"reports/helpers"
	"reports/services/queue"
	reportshandler "reports/services/reports-handler"
)

func main() {
	//initReceiver()
	reportshandler.Handle([]byte(`{"report":"admin-report", "option":["name", "email"]}`))
}

func initReceiver() {
	var conn = queue.QueueConnection

	defer conn.Close()

	ch, err := conn.Channel()

	helpers.LogError("Unable to connect to Channel", err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"reports", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	helpers.LogError("Unable to declare queue", err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	helpers.LogError("Unable to consume msg", err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			reportshandler.Handle(d.Body)

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
