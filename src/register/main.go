package main

import (
	"context"
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/taskworker"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/wflib"
)

func main() {
	config := common.NewConfigFromEnv()
	client, err := config.GetGrpcClient()
	if err != nil {
		log.Fatal(err)
	}

	// First, register the Task Definition
	log.Default().Print("Registering task definition")
	worker, err := taskworker.NewTaskWorker(config, src.Greet, "greet")
	if err != nil {
		log.Fatal(err)
	}
	err = worker.RegisterTaskDef(true)
	if err != nil {
		log.Fatal("Failed to register the task definition", err)
	}

	// Register the Workflow Spec
	workflow := wflib.NewWorkflow(src.QuickstartWorkflow, "quickstart")
	putWf, err := workflow.Compile()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := (*client).PutWfSpec(context.Background(), putWf)
	if err != nil {
		log.Fatal(err)
	}
	common.PrintProto(resp)
}
