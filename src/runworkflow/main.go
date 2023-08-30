package main

import (
	"context"
	"lh-quickstart-go/src"

	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common/model"
)

func main() {
	_, client := src.LoadConfigAndClient()
	name := "bill"

	common.PrintResp((*client).RunWf(
		context.Background(),
		&model.RunWfRequest{
			WfSpecName: "my-workflow",
			Variables: map[string]*model.VariableValue{
				"name": {
					Str:  &name,
					Type: model.VariableType_STR,
				},
			},
		}))
}
