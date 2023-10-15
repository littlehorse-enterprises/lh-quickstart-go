package main

import (
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/taskworker"
)

func main() {
	config := common.NewConfigFromEnv()

	worker, err := taskworker.NewTaskWorker(config, src.Greet, "greet")

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
