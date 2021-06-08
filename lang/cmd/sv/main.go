package main

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
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
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (stdrwc) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (stdrwc) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}

func run(ctx context.Context, args []string) error {
	cfg := zap.NewDevelopmentConfig()
	//cfg.OutputPaths = []string{
	//	"log.txt",
	//}
	logger, err := cfg.Build()
	if err != nil {
		log.Printf("failed to create logger: %v\n", err)
		os.Exit(exitCodeErr)
	}
	defer logger.Sync()
	logger.Info("Starting up...")

	stream := jsonrpc2.NewStream(stdrwc{})
	srv := systemverilog.NewServer()
	var conn jsonrpc2.Conn
	srv.Ctx, conn, srv.Client = protocol.NewServer(ctx, srv, stream, logger)

	select {
	case <-ctx.Done():
		logger.Info("Signal received")
		conn.Close()
	case <-conn.Done():
		logger.Info("Client disconnected")
	}

	logger.Info("Stopped...")
	return nil
}
