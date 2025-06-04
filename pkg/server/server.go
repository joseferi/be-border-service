package server

import (
	"be-border-service/internal/config"
	"be-border-service/internal/routes"

	"be-border-service/pkg/logger"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type httpServer struct {
	AppName string
	AppPort int
	router  routes.Router
	config  *config.Config
}

func NewHTTPServer(cfg *config.Config) Server {
	return &httpServer{
		AppName: cfg.Server.Name,
		AppPort: cfg.Server.Port,
		router:  routes.NewRouter(cfg),
		config:  cfg,
	}
}

// Run runs the http server gracefully
// returns:
//
//	err: error operation
func (h *httpServer) Run(ctx context.Context) error {
	var err error

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", h.AppPort),
		Handler: h.router.Route(),
		// ReadTimeout:  h.config.pp.ReadTimeout,
		// WriteTimeout: h.config.App.WriteTimeout,
	}

	go func() {
		err = server.ListenAndServe()
		if err != http.ErrServerClosed {
			slog.Error(fmt.Sprintf("http server got %v", err)) // temporary !
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed : %v", err)

	}

	slog.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

func (h *httpServer) Done() {
	logger.Info("server stopeed")
}
