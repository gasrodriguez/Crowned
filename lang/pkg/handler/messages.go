package handler

import (
	"context"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
)

func (o *Handler) LogError(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeError,
	})
	util.CheckError(err)
}

func (o *Handler) LogWarning(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeWarning,
	})
	util.CheckError(err)
}

func (o *Handler) LogInfo(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeInfo,
	})
	util.CheckError(err)
}

func (o *Handler) LogMessage(message string) {
	err := o.Client.LogMessage(context.TODO(), &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeLog,
	})
	util.CheckError(err)
}

func (o *Handler) ShowError(message string) {
	err := o.Client.ShowMessage(context.TODO(), &protocol.ShowMessageParams{
		Message: message,
		Type:    protocol.MessageTypeError,
	})
	util.CheckError(err)
}

func (o *Handler) ShowWarning(message string) {
	err := o.Client.ShowMessage(context.TODO(), &protocol.ShowMessageParams{
		Message: message,
		Type:    protocol.MessageTypeWarning,
	})
	util.CheckError(err)
}

func (o *Handler) ShowInfo(message string) {
	err := o.Client.ShowMessage(context.TODO(), &protocol.ShowMessageParams{
		Message: message,
		Type:    protocol.MessageTypeInfo,
	})
	util.CheckError(err)
}

func (o *Handler) ShowMessage(message string) {
	err := o.Client.ShowMessage(context.TODO(), &protocol.ShowMessageParams{
		Message: message,
		Type:    protocol.MessageTypeLog,
	})
	util.CheckError(err)
}
