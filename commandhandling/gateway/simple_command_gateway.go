package gateway

import (
	"axon-go/commandhandling"
	"context"
)

type SimpleCommandGateway struct {
	CommandBus commandhandling.CommandBus
}

func NewSimpleCommandGateway() *SimpleCommandGateway {
	defaultCommandBus := commandhandling.NewSimpleCommandBus()
	return NewSimpleCommandGatewayWithBus(defaultCommandBus)
}

func NewSimpleCommandGatewayWithBus(commandBus commandhandling.CommandBus) *SimpleCommandGateway {
	return &SimpleCommandGateway{
		CommandBus: commandBus,
	}
}

func (a *SimpleCommandGateway) Send(ctx context.Context, command interface{}) error {
	commandMessage := commandhandling.NewGenericCommandMessage(command)
	a.CommandBus.Dispatch(ctx, commandMessage)
	return nil
}

func (a *SimpleCommandGateway) SendWithCallback(ctx context.Context, command interface{},
	callback commandhandling.CommandCallback) error {
	commandMessage := commandhandling.NewGenericCommandMessage(command)
	a.CommandBus.DispatchWithCallback(ctx, commandMessage, callback)
	return nil
}

func (a *SimpleCommandGateway) SendAndWait(ctx context.Context, command interface{}) (result interface{}, err error) {
	futureCallback := commandhandling.NewFutureCallback()
	a.SendWithCallback(ctx, command, futureCallback)
	resp, err := futureCallback.GetResult()
	return resp, err
}

func (a *SimpleCommandGateway) SendAndWaitWithTimeout(ctx context.Context, command interface{},
	timeout int64) (result interface{}, err error) {
	// todo it is not really implementation, will do it in future
	return a.SendAndWait(ctx, command)
}
