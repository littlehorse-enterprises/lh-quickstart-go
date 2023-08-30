package src

import (
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common/model"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/wflib"
	"log"
)

func Greet(name string) string {
	log.Print("Executing task greet with input variable: ", name)
	if name == "obi-wan" {
		return "hello there"
	} else {
		return "hello, " + name
	}
}

func MyWorkflow(thread *wflib.ThreadBuilder) {
	nameVar := thread.AddVariableWithDefault("name", model.VariableType_STR, "Qui-Gon Jinn")

	// Make it searchable
	nameVar.WithIndex(model.IndexType_REMOTE_INDEX)

	thread.Execute("greet", nameVar)
}
