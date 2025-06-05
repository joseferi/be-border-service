package cmd

import (
	"be-border-service/cmd/migration"
	"be-border-service/cmd/server"
	workerserver "be-border-service/cmd/worker_server"
	"be-border-service/internal/config"
	"be-border-service/pkg/logger"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.RegisterConfiguration()
	logger.Setup(cfg.Server.Env)
	rootCmd := &cobra.Command{}

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		cancel()
	}()

	cmd := []*cobra.Command{
		{
			Use:   "http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				server.Start(ctx, &cfg)
			},
		},
		{
			Use:   "db:migrate",
			Short: "Database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.MigrationDatabase()
			},
		},
		{
			Use:   "run:workers:server",
			Short: "Runing background jobs ...",
			Run: func(cmd *cobra.Command, args []string) {
				workerserver.Start(&cfg)
			},
		},
	}
	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
