package src

import (
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common/model"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/wflib"
)

func QuickstartWorkflow(thread *wflib.WorkflowThread) {
	// Declare an input variable and make it searchable
	nameVar := thread.AddVariable("input-name", model.VariableType_STR).Searchable()

	// Execute a task and pass in the variable.
	thread.Execute("greet", nameVar)
}
