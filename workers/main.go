package main

import (
	workflow "lh-quickstart-go"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func main() {

	config := littlehorse.NewConfigFromEnv()

	verifyIdentityWorker, err1 := littlehorse.NewTaskWorker(config, workflow.VerifyIdentity, "verify-identity")
	notifyCustomerVerifiedWorker, err2 := littlehorse.NewTaskWorker(config, workflow.NotifyCustomerVerified, "notify-customer-verified")
	notifyCustomerNotVerifiedWorker, err3 := littlehorse.NewTaskWorker(config, workflow.NotifyCustomerNotVerified, "notify-customer-not-verified")

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal(err1, err2, err3)
	}

	defer func() {
		log.Default().Print("Shutting down task worker")
		verifyIdentityWorker.Close()
		notifyCustomerVerifiedWorker.Close()
		notifyCustomerNotVerifiedWorker.Close()
	}()

	log.Default().Print("Starting Task Worker")
	verifyIdentityWorker.Start()
	notifyCustomerVerifiedWorker.Start()
	notifyCustomerNotVerifiedWorker.Start()

}
