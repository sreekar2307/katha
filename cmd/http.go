package main

import (
	"context"
	"fmt"
	"github.com/sreekar2307/khata/controller/http"
	"github.com/sreekar2307/khata/controller/http/middleware"
	v1 "github.com/sreekar2307/khata/controller/http/v1"
	"log"
	"time"
)

func StartHTTPCmd(pCtx context.Context) error {
	d, err := GetDeps()
	if err != nil {
		return fmt.Errorf("failed to get dependencies: %w", err)
	}
	v1Controller := v1.NewV1Controller(d.ExpenseService, d.UserService, d.LedgerService)
	httpController := http.NewController(v1Controller)
	m := middleware.NewMiddleware(d.UserService)
	server, err := http.NewServer(d.Conf.Server, m, httpController)
	if err != nil {
		return fmt.Errorf("failed to create HTTP server: %w", err)
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()
	<-pCtx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
