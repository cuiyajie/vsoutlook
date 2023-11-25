package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/subcommands"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/redis"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/router"
)

type serveCmd struct{}

func (cmd *serveCmd) Name() string {
	return "serve"
}

func (cmd *serveCmd) Synopsis() string {
	return "runs api server"
}

func (cmd *serveCmd) Usage() string {
	return "run with: ./server serve\n"
}

func (cmd *serveCmd) SetFlags(fs *flag.FlagSet) {}

func (cmd *serveCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	// go site.startSiteApp()

	redis.ConnectRedis()
	models.ReadyDB()

	startServer()

	return subcommands.ExitSuccess
}

func startServer() {
	log.Printf("server start")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addrFmt := "0.0.0.0:%v"
	if config.Dev {
		addrFmt = "localhost:%v"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(addrFmt, config.Get("HTTP_PORT")),
		Handler: router.SetupRouter(),
	}

	go func() {
		fmt.Printf("Listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	// log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	fmt.Println("Server exits")
}
