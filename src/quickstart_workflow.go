package src

import (
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/common/model"
	"github.com/littlehorse-enterprises/littlehorse/sdk-go/wflib"
)

func QuickstartWorkflow(thread *wflib.WorkflowThread) {
	// Declare an input variable
	nameVar := thread.AddVariable("input-name", model.VariableType_STR)

	// Make the variable searchable
	nameVar.WithIndex(model.IndexType_REMOTE_INDEX).Persistent()

	// Execute a task and pass in the variable.
	thread.Execute("greet", nameVar)
}
