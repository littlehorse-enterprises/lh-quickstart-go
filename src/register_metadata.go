package src

import (
	"context"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/lhproto"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"

	"lh-quickstart-go/src"
)

const (
	WORKFLOW_NAME                     = "quickstart"
	IDENTITY_VERIFIED_EVENT           = "identity-verified"
	VERIFY_IDENTITY_TASK              = "verify-identity"
	NOTIFY_CUSTOMER_VERIFIED_TASK     = "notify-customer-verified"
	NOTIFY_CUSTOMER_NOT_VERIFIED_TASK = "notify-customer-not-verified"
)

func QuickstartWorkflow(wf *littlehorse.WorkflowThread) {
	// Declare the input variables for the workflow.
	firstName := wf.DeclareStr("first-name").Searchable().Required()
	lastName := wf.DeclareStr("last-name").Searchable().Required()

	// Social Security Numbers are sensitive, so we mask the variable.
	ssn := wf.DeclareInt("ssn").MaskedValue().Required()

	// Internal variable representing whether the customer's identity has been verified.
	identityVerified := wf.DeclareBool("identity-verified").Searchable()

	// Call the verify-identity task and retry it up to 3 times if it fails
	wf.Execute(VERIFY_IDENTITY_TASK, firstName, lastName, ssn).WithRetries(3)

	// Make the WfRun wait until the event is posted or if the timeout is reached
	identityVerificationResult := wf.WaitForEvent(IDENTITY_VERIFIED_EVENT).Timeout(60 * 60 * 24 * 3)

	exceptionName := littlehorse.Timeout

	wf.HandleError(&identityVerificationResult, &exceptionName, func(handler *littlehorse.WorkflowThread) {
		handler.Execute(NOTIFY_CUSTOMER_NOT_VERIFIED_TASK, firstName, lastName)
		message := "Unable to verify customer identity in time."
		handler.Fail(nil, "customer-not-verified", &message)
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

func main() {
	config := littlehorse.NewConfigFromEnv()
	client, err := config.GetGrpcClient()
	if err != nil {
		log.Fatal(err)
	}

	// First, register the TaskDef
	log.Default().Print("Registering TaskDefs")
	verifyIdentityWorker, err := littlehorse.NewTaskWorker(config, src.VerifyIdentity, "verify-identity")
	if err != nil {
		log.Fatal("Failed to create verify-identity worker", err)
	}
	notifyCustomerVerifiedWorker, err2 := littlehorse.NewTaskWorker(config, src.NotifyCustomerVerified, "notify-customer-verified")
	if err2 != nil {
		log.Fatal("Failed to create notify-customer-verified worker", err2)
	}
	notifyCustomerNotVerifiedWorker, err3 := littlehorse.NewTaskWorker(config, src.NotifyCustomerNotVerified, "notify-customer-not-verified")
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
