package main

import (
	"context"
	"os"
	"os/signal"
	"workplaces/server/internals/app"
	"workplaces/server/internals/cfg"

	log "github.com/sirupsen/logrus"
)

func main() {
	config := cfg.LoadConfig()
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		server.Shutdown()
		cancel()
	}()

	server.Serve()
}
