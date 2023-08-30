package main

import (
	"context"
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/wflib"
)

func main() {
	_, client := src.LoadConfigAndClient()
	workflow := wflib.NewWorkflow(src.MyWorkflow, "my-workflow")
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
