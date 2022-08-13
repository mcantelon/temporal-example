package workflows

import (
	"go.temporal.io/sdk/workflow"
)

// ExampleParentWorkflow is a Workflow Definition
//
// This Workflow Definition demonstrates how to start a Child Workflow Execution from a Parent Workflow Execution.
//
// Each Child Workflow Execution starts a new Run.
// The Parent Workflow Execution is notified only after the completion of last Run of the Child Workflow Execution.
func ExampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	currentState := "started" // This could be any serializable struct.
	queryType := "current_state"
	err := workflow.SetQueryHandler(ctx, queryType, func() (string, error) {
		return currentState, nil
	})
	if err != nil {
		currentState = "failed to register query handler"
		return "", err
	}

	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: "EXAMPLE-CHILD-WORKFLOW-ID",
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var result string
	err = workflow.ExecuteChildWorkflow(ctx, ExampleChildWorkflow, "Some parameters").Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return "", err
	}

	logger.Info("Parent execution completed.", "Result", result)
	return result, nil
}
