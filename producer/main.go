package main

import (
	"log"
	mb "producer/kafka"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var kafkaCfg = mb.KafkaCfg{
	Server:   "localhost:29092",
	ClientID: "TestProject",
	Acks:     "all",
}
var topic = "TestTopic"

func main() {

	kafkaProducer, err := mb.KafkCli(kafkaCfg)
	if err != nil {
		log.Fatalf("Fatal to create Mb: %s", err)
	}

	app := fiber.New()
	count := 0
	app.Post("/hello:name", func(c *fiber.Ctx) error {

		message := c.Params("name") + strconv.Itoa(count)

		status, err := kafkaProducer.SendMessage([]byte(message), topic)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString(status)
	})

	app.Listen(":8080")

}
