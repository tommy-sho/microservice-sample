package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tommy-sho/microservice-sample/cmd/user/app"
)

func main() {
	option := &app.Option{
		Port:                 "8000",
		ConnectTimeout:       100,
		KeepAlive:            10,
		GracefulPeriod:       10,
		MetricsAdnHealthPort: "9000",
		LogLevel:             "Info",
	}

	server := app.NewServer(option)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	)

	go func() {
		<-stopChan
		log.Println("start server shutdown")
		server.Shutdown(time.Second)
	}()

	if err := server.Serve(); err != nil {
		log.Fatalf("failed to serve: %#v", err)
	}
}
