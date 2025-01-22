package main

import (
	"context"
	workflow "lh-quickstart-go"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/lhproto"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func main() {
	config := littlehorse.NewConfigFromEnv()
	client, err := config.GetGrpcClient()
	if err != nil {
		log.Fatal(err)
	}

	(*client).PutExternalEventDef(context.Background(),
		&lhproto.PutExternalEventDefRequest{
			Name: workflow.IDENTITY_VERIFIED_EVENT,
		},
	)

	// First, register the TaskDef
	log.Default().Print("Registering TaskDefs")
	verifyIdentityWorker, err := littlehorse.NewTaskWorker(config, workflow.VerifyIdentity, "verify-identity")
	if err != nil {
		log.Fatal("Failed to create verify-identity worker", err)
	}
	notifyCustomerVerifiedWorker, err2 := littlehorse.NewTaskWorker(config, workflow.NotifyCustomerVerified, "notify-customer-verified")
	if err2 != nil {
		log.Fatal("Failed to create notify-customer-verified worker", err2)
	}
	notifyCustomerNotVerifiedWorker, err3 := littlehorse.NewTaskWorker(config, workflow.NotifyCustomerNotVerified, "notify-customer-not-verified")
	if err3 != nil {
		log.Fatal("Failed to create notify-customer-not-verified worker", err3)
	}

	err = verifyIdentityWorker.RegisterTaskDef()
	if err != nil {
		log.Fatal("Failed to register verify-identity", err)
	}
	err = notifyCustomerVerifiedWorker.RegisterTaskDef()
	if err != nil {
		log.Fatal("Failed to register notify-customer-verified", err)
	}
	err = notifyCustomerNotVerifiedWorker.RegisterTaskDef()
	if err != nil {
		log.Fatal("Failed to register notify-customer-not-verified", err)
	}

	// Register the Workflow Spec
	log.Default().Print("Registering WfSpec")
	workflow := littlehorse.NewWorkflow(workflow.QuickstartWorkflow, "quickstart")
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
