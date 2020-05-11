package queue

import (
	"log"
)

func init() {
	// amqp://guest:guest@localhost:5672/
	log.Println("Initialized QUEUE URL")
	// fmt.Sprintf("amqp://%v:%v@%v/",
	// 	Config.QueueUser,
	// 	Config.QueuePass,
	// 	Config.QueueHost,
	// )

}
