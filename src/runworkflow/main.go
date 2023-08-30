package main

import (
	"context"
	"lh-quickstart-go/src"
	"log"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common/model"
)

func main() {
	_, client := src.LoadConfigAndClient()
	name := "bill"
	wfId, err := (*client).RunWf(
		context.Background(),
		&model.RunWfRequest{
			WfSpecName: "my-workflow",
			Variables: map[string]*model.VariableValue{
				"name": {
					Str:  &name,
					Type: model.VariableType_STR,
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("got wfRunModel Id:", wfId)

}
