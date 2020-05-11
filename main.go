package main

import (
	"fmt"
	"log"
	"reports/config"
	"reports/helpers"
	"reports/routes"
	"reports/services/database"
	"reports/services/queue"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	// load configurations
	serverConfig := config.Config

	//  initialize all configurations
	// config.InitDB(serverConfig)
	// defer config.DB.Close()

	gin.SetMode(serverConfig.RunMode)
	//gin.DisableConsoleColor()

	initReceiver()

	router := routes.Router()

	router.Run()

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

	// var info map[string]interface{}
	// var pool = map[float64]string{
	// 	0: "9",
	// 	7: "1",
	// }

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			rows, err := database.DB.Raw("select * from users;").Rows()
			helpers.LogError("Error in runnig query", err)
			res := helpers.MapScan(rows)
			fmt.Print(res)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
