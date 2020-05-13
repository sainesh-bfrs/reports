package main

/*
 * File: main.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"log"
	"reports/config"
	"reports/helpers"
	"reports/services/queue"
	reportshandler "reports/services/reports-handler"

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

	// router := routes.Router()

	// router.Run()

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
			// rows, err := database.DB.Raw("select id, company_id, first_name, last_name from users;").Rows()
			// helpers.LogError("Error in runnig query", err)
			// res := helpers.MapScan(rows)

			// data := helpers.PrepareCSVData(res)

			// helpers.WriteCSV(data, "storage/data.csv")

			// url := aws.Upload("storage/data.csv", "test/data.csv")

			// mailer := mail.Mailer{
			// 	To:      []string{"sainesh.mamgain@kartrocket.com"},
			// 	Subject: "Test Mail",
			// 	Body:    "URL for file: " + url,
			// }

			// mailer.Send()

			reportshandler.Handle(d.Body)

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
