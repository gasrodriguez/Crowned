package main

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"github.com/gasrodriguez/crowned/pkg/server"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(server.ExitCodeInterrupt)
	}()
	svServer := systemverilog.NewServer()
	if err := svServer.Run(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(server.ExitCodeErr)
	}
}
