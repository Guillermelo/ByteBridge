package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"ByteBridge/internals/dispatcher"
	"ByteBridge/internals/serverconn"
)

const (
	version          = "1.0.0"
	MaxDevicesAmount = 10
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	fmt.Println("Starting ByteBridge", version)
	var cfg config

	flag.IntVar(&cfg.port, "port", 6000, "api port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|production|staging)")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ByteBridge := &application{
		config: cfg,
		logger: log,
	}

	Addr := fmt.Sprintf(":%d", cfg.port)
	ByteBridge.logger.Info("ByteBridge Started at ", "addr", Addr)
	// Worker := &dispatcher.Worker{
	// 	ID: "1",
	// }
	// Worker.Work()
	Dispatcher := dispatcher.Dispatcher{
		ConnPool: make(chan *serverconn.ServerConn, MaxDevicesAmount),
	}
	go Dispatcher.Dispatch()
	serverconn.FillConnPool(Addr, Dispatcher.ConnPool)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
