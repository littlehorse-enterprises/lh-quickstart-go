package workflow

import (
	"errors"
	"math/rand"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func VerifyIdentity(firstName, lastName string, ssn int) (string, error) {
	if rand.Float32() < 0.25 {
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

const (
	WORKFLOW_NAME                     = "quickstart"
	IDENTITY_VERIFIED_EVENT           = "identity-verified"
	VERIFY_IDENTITY_TASK              = "verify-identity"
	NOTIFY_CUSTOMER_VERIFIED_TASK     = "notify-customer-verified"
	NOTIFY_CUSTOMER_NOT_VERIFIED_TASK = "notify-customer-not-verified"
)

func QuickstartWorkflow(wf *littlehorse.WorkflowThread) {
	// Declare the input variables for the workflow.
	firstName := wf.DeclareStr("first-name", nil).Searchable().Required()
	lastName := wf.DeclareStr("last-name", nil).Searchable().Required()

	// Social Security Numbers are sensitive, so we mask the variable.
	ssn := wf.DeclareInt("ssn", nil).MaskedValue().Required()

	// Internal variable representing whether the customer's identity has been verified.
	identityVerified := wf.DeclareBool("identity-verified", nil).Searchable()

	// Call the verify-identity task and retry it up to 3 times if it fails
	wf.Execute(VERIFY_IDENTITY_TASK, firstName, lastName, ssn).WithRetries(3)

	// Make the WfRun wait until the event is posted or if the timeout is reached
	identityVerificationResult := wf.WaitForEvent(IDENTITY_VERIFIED_EVENT).Timeout(60 * 60 * 24 * 3)
	exceptionName := littlehorse.Timeout
	wf.HandleError(identityVerificationResult, &exceptionName, func(handler *littlehorse.WorkflowThread) {
		handler.Execute(NOTIFY_CUSTOMER_NOT_VERIFIED_TASK, firstName, lastName)
		message := "Unable to verify customer identity in time."
		handler.Fail("d", "customer-not-verified", &message)
	})

	// Assign the output of the ExternalEvent to the `identityVerified` variable.
	identityVerified.Assign(identityVerificationResult)

	// Notify the customer if their identity was verified or not
	wf.DoIfElse(
		identityVerified.IsEqualTo(true),
		func(ifBody *littlehorse.WorkflowThread) {
			ifBody.Execute(NOTIFY_CUSTOMER_VERIFIED_TASK, firstName, lastName)
		},
		func(elseBody *littlehorse.WorkflowThread) {
			elseBody.Execute(NOTIFY_CUSTOMER_NOT_VERIFIED_TASK, firstName, lastName)
		},
	)
}
