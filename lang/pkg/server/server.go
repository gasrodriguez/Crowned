package server

import (
	"context"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
	"log"
	"os"
)

const (
	ExitCodeErr       = 1
	ExitCodeInterrupt = 2
)

type Server struct {
	protocol.Server
	Client protocol.Client
}

func (o *Server) Run(ctx context.Context, args []string) error {
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
	defer o.loggerSync(logger)
	logger.Info("Starting up...")

	stream := jsonrpc2.NewStream(Stdinout{})
	ctx, conn, client := protocol.NewServer(ctx, o, stream, logger)
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
	return nil
}

func (o *Server) loggerSync(logger *zap.Logger) {
	err := logger.Sync()
	util.CheckError(err)
}

func (o *Server) LogMessage(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeLog,
	})
	util.CheckError(err)
}

func (o *Server) LogError(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeError,
	})
	util.CheckError(err)
}
