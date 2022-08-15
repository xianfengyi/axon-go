package commandhandling

import (
	"context"
	"errors"
	"sync"
)

// ErrHandlerAlreadySet is when a handler is already registered for a command.
var ErrHandlerAlreadySet = errors.New("handler is already set")

// ErrHandlerNotFound is when no handler can be found.
var ErrHandlerNotFound = errors.New("no handlers for command")

var DefaultCallback CommandCallback = LoggingCallback{}

// SimpleCommandBus Implementation of the CommandBus that dispatches commands to the handlers subscribed to that specific command's name.
type SimpleCommandBus struct {
	handlers   map[string]CommandHandler
	handlersMu sync.RWMutex
}

// NewSimpleCommandBus creates a SimpleCommandBus.
func NewSimpleCommandBus() *SimpleCommandBus {
	return &SimpleCommandBus{
		handlers:   make(map[string]CommandHandler),
		handlersMu: sync.RWMutex{},
	}
}

func (bus *SimpleCommandBus) Dispatch(ctx context.Context, commandMessage CommandMessage) error {
	err := bus.doDispatch(ctx, commandMessage, DefaultCallback)
	return err
}

func (bus *SimpleCommandBus) DispatchWithCallback(ctx context.Context, commandMessage CommandMessage, callback CommandCallback) error {
	err := bus.doDispatch(ctx, commandMessage, callback)
	return err
}

// Subscribe the given {@code handler} to commands with given {@code commandName}. If a subscription already
// exists for the given name, the configured will return error ErrHandlerAlreadySet
func (bus *SimpleCommandBus) Subscribe(commandName string, handler CommandHandler) error {
	bus.handlersMu.Lock()
	defer bus.handlersMu.Unlock()

	if _, ok := bus.handlers[commandName]; ok {
		return ErrHandlerAlreadySet
	}
	bus.handlers[commandName] = handler
	return nil
}

// Performs the actual dispatching logic.
func (bus *SimpleCommandBus) doDispatch(ctx context.Context, command CommandMessage, callback CommandCallback) error {
	handler, err := bus.findCommandHandlerFor(command)
	if err == nil {
		bus.handle(ctx, command, handler, callback)
	}
	return nil
}

func (bus *SimpleCommandBus) findCommandHandlerFor(message CommandMessage) (CommandHandler, error) {
	if _, ok := bus.handlers[message.GetCommandName()]; ok {
		return bus.handlers[message.GetCommandName()], nil
	}
	return nil, ErrHandlerNotFound
}

func (bus *SimpleCommandBus) handle(ctx context.Context, command CommandMessage,
	handler CommandHandler, commandCallback CommandCallback) {
	// execute handle command
	resp, err := handler.HandleCommand(ctx, command.GetPayload())
	commandResultMessage := NewGenericCommandResultMessage(resp, err)
	//  callback get result
	commandCallback.OnResult(command, commandResultMessage)
}
