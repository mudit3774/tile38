package server

import (
	"errors"
	"github.com/tidwall/tile38/internal/log"
)

var errCannotExecuteServerHook = errors.New("cannot execute server hook")

type SHooks interface {
	PreStartHook() error
	PreCmdHook(cmd Message) error
}

type NoOpHook struct {
}

func NewNoOpHook(server Server) SHooks {
	return NoOpHook{}
}

func (n NoOpHook) PreStartHook() error {
	log.Debug("Executing pre-start hook")
	return nil
}

func (n NoOpHook) PreCmdHook(cmd Message) error {
	log.Debugf("Executing pre-cmd hook for cmd %s", cmd._command)
	return nil
}

