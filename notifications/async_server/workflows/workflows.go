package workflows

import (
	"time"

	"github.com/ahsansaif47/home-kitchens/notifications/async_server/activities"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func SendOTPWorkflow(ctx workflow.Context, receipient, otp string) error {

	actOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 10,
			MaximumAttempts:    3,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, actOptions)

	err := workflow.ExecuteActivity(ctx, activities.SendOTPActivity, receipient, otp).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to send OTP email", "error", err)
		return err
	}

	workflow.GetLogger(ctx).Error("OTP email sent successfully", "receipient", receipient)
	return nil
}
