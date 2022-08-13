package activities

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

func ExampleActivity(ctx context.Context, value string) (string, error) {
	activity.GetLogger(ctx).Info("ExampleActivity called.", "Value", value)
	time.Sleep(5 * time.Second)
	return "Processed: " + value, nil
}
