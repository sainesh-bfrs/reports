package queue

import (
	"fmt"
	"log"
	"reports/config"
	"reports/helpers"

	"github.com/streadway/amqp"
)

// QueueConnection ...
var QueueConnection *amqp.Connection

func queueConnectionString() string {
	return fmt.Sprintf("amqp://%v:%v@%v/",
		config.Config.QueueUser,
		config.Config.QueuePass,
		config.Config.QueueHost,
	)
}

func init() {
	log.Println("Initialized QUEUE Connection")

	var err error
	QueueConnection, err = amqp.Dial(queueConnectionString())

	helpers.LogError("Unable to connect to queue", err)
}
