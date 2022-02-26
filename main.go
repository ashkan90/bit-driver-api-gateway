package main

import (
	"context"
	"flag"
	"github.com/ashkan90/bit-driver-api-gateway/proxy/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashkan90/bit-driver-api-gateway/proxy/router"
)

func main() {
	var proxyConfigPath string
	var proxyConfig = &config.ServiceConfig{}
	var logger = log.Default()

	flag.StringVar(&proxyConfigPath, "proxy-services", "", "")
	flag.Parse()

	if proxyConfigPath == "" {
		panic("Proxy services has not been set.")
	}

	proxyConfig.ImportInto(proxyConfigPath)

	var pRouter = &router.ProxyRouter{
		Logger:  logger,
		Config: proxyConfig,
	}

	var server = &http.Server{
		Addr:    ":4050",
		Handler: pRouter.NewRouter(),
	}

	go func() {
		logger.Fatal(server.ListenAndServe())
	}()

	gracefulShutdown(logger, server)
}

func gracefulShutdown(logger *log.Logger, s *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
