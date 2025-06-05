package workerserver

import (
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"be-border-service/internal/handler/tasks"
	"be-border-service/pkg/logger"
	"be-border-service/pkg/workerx"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Start(config *config.Config) {

	srv := workerx.NewAsynqserver(config)

	// Register handlers
	srv.Register(constants.TaskHealthCheck, tasks.HealthCheckHandler())

	// Graceful shutdown support
	logger.Info("[🚀] Asynq Server starting ....")
	go func() {
		if err := srv.Run(); err != nil {
			logger.Fatal(fmt.Sprintf("Asynq server error :%v", err), logger.EventName(constants.WorkerStarting))
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Info("[Asynq Server] Shutting down...")
}
