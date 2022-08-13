package http

import (
	"fmt"
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"

	"github.com/mcantelon/temporal-example/workflows"
)

func getTemplateContent(urlPath string) string {
	filePath := getServerDirectory() + "/template" + urlPath + ".html"

	var content []byte
	var err error

	if checkFileExists(filePath) {
		fmt.Println("Returning " + filePath)

		content, err = ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Print(err)
		}
	}

	return string(content)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func getServerDirectory() string {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename)
}

func startWorkflow() (string, error) {
	// Attempt to start workflow
        workflowID := "parent-workflow_" + uuid.New()
        workflowOptions := client.StartWorkflowOptions{
                ID:        workflowID,
                TaskQueue: "example-task-queue",
        }

        workflowRun, err := appState.client.ExecuteWorkflow(context.Background(), workflowOptions, workflows.ExampleParentWorkflow)
        if err != nil {
		return "Unable to execute workflow", err
        }

	message := "Workflow ID " +  workflowRun.GetID() + ", run ID " + workflowRun.GetRunID()

	return message, nil
}
