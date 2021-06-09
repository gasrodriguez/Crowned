package lsp

import (
	"context"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

const (
	ExitCodeErr       = 1
	ExitCodeInterrupt = 2
)

type Handler struct {
	protocol.Server
	Client protocol.Client
}

func (o *Handler) Run(server protocol.Server, args []string) {
	// ToDo: use server options
	cfg := zap.NewDevelopmentConfig()
	//cfg.OutputPaths = []string{
	//	"log.txt",
	//}
	logger, err := cfg.Build()
	if err != nil {
		log.Printf("failed to create logger: %v\n", err)
		os.Exit(ExitCodeErr)
	}
	defer loggerSync(logger)
	logger.Info("Starting up...")

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancelFunc()
	}()

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancelFunc()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(ExitCodeInterrupt)
	}()

	stream := jsonrpc2.NewStream(util.Stdinout{})
	ctx, conn, client := protocol.NewServer(ctx, server, stream, logger)
	o.Client = client

	select {
	case <-ctx.Done():
		logger.Info("Signal received")
		err := conn.Close()
		util.CheckError(err)
	case <-conn.Done():
		logger.Info("Client disconnected")
	}

	logger.Info("Stopped...")
}

func loggerSync(logger *zap.Logger) {
	err := logger.Sync()
	util.CheckError(err)
}
