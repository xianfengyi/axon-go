package commandhandling

import (
	"context"
	"sync"
)

type AsynchronousCommandBus struct {
	SimpleCommandBus
	resp chan response
}

type response struct {
	data interface{}
	err  error
}

func NewAsynchronousCommandBus() *AsynchronousCommandBus {
	return &AsynchronousCommandBus{
		SimpleCommandBus: SimpleCommandBus{
			handlers:   make(map[string]CommandHandler),
			handlersMu: sync.RWMutex{},
		},
		resp: make(chan response),
	}
}

func (bus *AsynchronousCommandBus) Dispatch(ctx context.Context, commandMessage CommandMessage) error {
	err := bus.doDispatch(ctx, commandMessage, DefaultCallback)
	return err
}

func (bus *AsynchronousCommandBus) DispatchWithCallback(ctx context.Context, commandMessage CommandMessage, callback CommandCallback) error {
	err := bus.doDispatch(ctx, commandMessage, callback)
	return err
}

// Performs the actual dispatching logic.
func (bus *AsynchronousCommandBus) doDispatch(ctx context.Context, command CommandMessage, callback CommandCallback) error {
	handler, err := bus.findCommandHandlerFor(command)
	if err == nil {
		bus.handle(ctx, command, handler, callback)
	}
	return nil
}

func (bus *AsynchronousCommandBus) findCommandHandlerFor(message CommandMessage) (CommandHandler, error) {
	if _, ok := bus.handlers[message.GetCommandName()]; ok {
		return bus.handlers[message.GetCommandName()], nil
	}
	return nil, ErrHandlerNotFound
}

func (a *AsynchronousCommandBus) handle(ctx context.Context, command CommandMessage, handler CommandHandler,
	commandCallback CommandCallback) {
	// execute handle command with a new goroutines
	go func() {
		resp, err := handler.HandleCommand(ctx, command.GetPayload())
		ret := response{
			data: resp,
			err:  err,
		}
		a.resp <- ret
	}()
	out := <-a.resp
	commandResultMessage := NewGenericCommandResultMessage(out.data, out.err)
	//  callback get result
	commandCallback.OnResult(command, commandResultMessage)
}
