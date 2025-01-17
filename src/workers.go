package src

import (
	"errors"
	"log"
	"math/rand"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func VerifyIdentity(firstName, lastName string, ssn int) (string, error) {
	if rand.Float64() < 0.25 {
		return "", errors.New("the external identity verification API is down")
	}
	return "Successfully called external API to request verification for " + firstName + " " + lastName, nil
}

func NotifyCustomerVerified(firstName, lastName string) string {
	return "Notification sent to customer " + firstName + " " + lastName + " that their identity has been verified"
}

func NotifyCustomerNotVerified(firstName, lastName string) string {
	return "Notification sent to customer " + firstName + " " + lastName + " that their identity has not been verified"
}

func main() {

	config := littlehorse.NewConfigFromEnv()

	verifyIdentityWorker, err1 := littlehorse.NewTaskWorker(config, VerifyIdentity, "verify-identity")
	notifyCustomerVerifiedWorker, err2 := littlehorse.NewTaskWorker(config, NotifyCustomerVerified, "notify-customer-verified")
	notifyCustomerNotVerifiedWorker, err3 := littlehorse.NewTaskWorker(config, NotifyCustomerNotVerified, "notify-customer-not-verified")

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
