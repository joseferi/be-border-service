package cmd

import (
	"be-border-service/cmd/migration"
	"be-border-service/cmd/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func Start() {
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
				server.Start(ctx)
			},
		},
		{
			Use:   "db:migrate",
			Short: "Database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.MigrationDatabase()
			},
		},
	}
	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
