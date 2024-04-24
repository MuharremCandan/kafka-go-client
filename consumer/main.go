package main

import (
	mb "consumer/kafka"
	"log"
)

var kafkaCfg = mb.KafkaCfg{
	Server:  "localhost:29092",
	GroupID: "TestProject",
}
var topic = "TestTopic"

func main() {
	kafkaCli, err := mb.KafkCli(kafkaCfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatalf("Error to read message: %s", kafkaCli.ReadMessage(topic))

}
