package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/mcantelon/temporal-example/activities"
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

	w := worker.New(c, "example-task-queue", worker.Options{})

	w.RegisterWorkflow(workflows.ExampleParentWorkflow)
	w.RegisterWorkflow(workflows.ExampleChildWorkflow)
	w.RegisterActivity(activities.ExampleActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
