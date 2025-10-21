package asyncserver

import (
	"log"

	"github.com/ahsansaif47/home-kitchens/notifications/async_server/activities"
	"github.com/ahsansaif47/home-kitchens/notifications/async_server/workflows"
	"go.temporal.io/sdk/worker"
)

// TODO: Return error from all workers using channels

func SetupOTPWorker() error {
	w := worker.New(TClient, "otp-queue", worker.Options{})

	w.RegisterWorkflow(workflows.SendOTPWorkflow)
	w.RegisterActivity(activities.SendOTPActivity)

	var err error

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("unable to start worker: %v", err)
		}
	}()

	return err
}
