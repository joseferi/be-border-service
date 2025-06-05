package server

import (
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"be-border-service/pkg/logger"
	"be-border-service/pkg/server"
	"context"
	"fmt"
)

func Start(ctx context.Context, config *config.Config) {

	serve := server.NewHTTPServer(config)

	defer serve.Done()
	logger.Info(fmt.Sprintf("starting [%s] services ... %d", config.Server.Name, config.Server.Port),
		logger.EventName(constants.LogEventNameServiceStarting),
	)

	if err := serve.Run(ctx); err != nil {
		logger.Warn(fmt.Sprintf("service stopped, err : %s", err.Error()), logger.EventName(constants.LogEventNameServiceStarting))
	}

}
