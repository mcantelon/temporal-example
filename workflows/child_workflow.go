package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/mcantelon/temporal-example/activities"
)

// ExampleChildWorkflow is a Workflow Definition
func ExampleChildWorkflow(ctx workflow.Context, data string) (string, error) {
	logger := workflow.GetLogger(ctx)

        // RetryPolicy specifies how to automatically handle retries if an Activity fails.
        retrypolicy := &temporal.RetryPolicy{
                InitialInterval:    time.Second,
                BackoffCoefficient: 2.0,
                MaximumInterval:    time.Minute,
                MaximumAttempts:    3,
        }

        options := workflow.ActivityOptions{
                // Timeout options specify when to automatically timeout Activity functions.
                StartToCloseTimeout: time.Minute,
                // Optionally provide a customized RetryPolicy.
                // Temporal retries failures by default, this is just an example.
                RetryPolicy: retrypolicy,
        }
	ctx = workflow.WithActivityOptions(ctx, options)
	err := workflow.ExecuteActivity(ctx, activities.ExampleActivity, "some data").Get(ctx, nil)
	if err != nil {
		return "", err
	}

	results := "Child workflow processed: " + data + "."
	logger.Info("Child workflow execution started with data: " + data)

	return results, nil
}
