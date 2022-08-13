package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"

	"github.com/mcantelon/temporal-example/workflows"
)

func main() {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This Workflow ID can be a user supplied business logic identifier.
	workflowID := "parent-workflow_" + uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "example-task-queue",
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.ExampleParentWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	// Query the workflow we've just started
	queryType := "current_state"
	resp, err := c.QueryWorkflow(context.Background(), workflowID, "", queryType)
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}
	var query_result interface{}
	if err := resp.Get(&query_result); err != nil {
		log.Fatalln("Unable to decode query result", err)
	}
	log.Println("Received query result:", query_result)

	// Synchronously wait for the Workflow Execution to complete.
	// Behind the scenes the SDK performs a long poll operation.
	// If you need to wait for the Workflow Execution to complete from another process use
	// Client.GetWorkflow API to get an instance of the WorkflowRun.
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Failure getting workflow result", err)
	}
	log.Printf("Workflow result: %v", result)
}
