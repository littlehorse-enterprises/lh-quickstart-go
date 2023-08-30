package main

import (
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/taskworker"
)

func main() {
	config, _ := src.LoadConfigAndClient()

	tw, err := taskworker.NewTaskWorker(config, src.Greet, "greet")

	if err != nil {
		log.Fatal(err)
	}

	// Create the TaskDef
	err = tw.RegisterTaskDef(true)
	if err != nil {
		log.Fatal("Failed to register the task definition", err)
	}

	defer func() {
		log.Default().Print("Shutting down task worker")
		tw.Close()
	}()

	log.Default().Print("Starting Task Worker")
	tw.Start()

}
