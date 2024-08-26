package src

import (
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/lhproto"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/littlehorse"
)

func QuickstartWorkflow(thread *littlehorse.WorkflowThread) {
	// Declare an input variable and make it searchable
	nameVar := thread.AddVariable("input-name", lhproto.VariableType_STR).Searchable()

	// Execute a task and pass in the variable.
	thread.Execute("greet", nameVar)
}
