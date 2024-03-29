package main

import (
	"context"
	"flag"
	"github.com/ashkan90/bit-driver-api-gateway/config"
	"github.com/ashkan90/bit-driver-api-gateway/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var reset = "\033[0m"
var yellow = "\033[33m"

func main() {
	var proxyConfigPath string
	var logger = log.Default()
	logger.SetPrefix(yellow + "[INFO] " + reset)

	flag.StringVar(&proxyConfigPath, "proxy-services", "", "")
	flag.Parse()

	if proxyConfigPath == "" {
		panic("Proxy services has not been set.")
	}

	var proxyConfig, err = config.NewConfig(proxyConfigPath)
	if err != nil {
		panic(err)
	}

	var pRouter = &router.ProxyRouter{
		Logger: logger,
		Config: &proxyConfig,
	}

	var addr = os.Getenv("PORT")
	if addr != "" {
		proxyConfig.Server.Port = addr
	}
	var server = &http.Server{
		Addr:    ":" + proxyConfig.Server.Port,
		Handler: pRouter.NewRouter(),
	}

	log.Println("Server started at ", proxyConfig.Server.Port)

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
