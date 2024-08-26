package main

import (
	"context"
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func main() {
	config := littlehorse.NewConfigFromEnv()
	client, err := config.GetGrpcClient()
	if err != nil {
		log.Fatal(err)
	}

	// First, register the Task Definition
	log.Default().Print("Registering task definition")
	worker, err := littlehorse.NewTaskWorker(config, src.Greet, "greet")
	if err != nil {
		log.Fatal(err)
	}
	err = worker.RegisterTaskDef()
	if err != nil {
		log.Fatal("Failed to register the task definition", err)
	}

	// Register the Workflow Spec
	workflow := littlehorse.NewWorkflow(src.QuickstartWorkflow, "quickstart")
	putWf, err := workflow.Compile()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := (*client).PutWfSpec(context.Background(), putWf)
	if err != nil {
		log.Fatal(err)
	}
	littlehorse.PrintProto(resp)
}
