package main

import (
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func main() {
	config := littlehorse.NewConfigFromEnv()

	worker, err := littlehorse.NewTaskWorker(config, src.Greet, "greet")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		log.Default().Print("Shutting down task worker")
		worker.Close()
	}()

	log.Default().Print("Starting Task Worker")
	worker.Start()

}
