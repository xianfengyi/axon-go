package gateway

import (
	"axon-go/commandhandling"
	"context"
)

type CommandGateway interface {
	// Send the given command
	Send(ctx context.Context, Command interface{}) error

	// SendWithCallback Sends the given {@code command}, and have the result of the command's execution reported to the given
	//  {@code callback}.
	SendWithCallback(ctx context.Context, Command interface{}, callback commandhandling.CommandCallback) error

	// SendAndWait Sends the given command and wait for it to execute. The result of the execution is returned when available.
	// This method will block indefinitely ,until a result is available,or until the Thread is interrupted. When the
	// timeout is reached or the thread is interrupted, this method returns nil and the detail error.
	SendAndWait(ctx context.Context, Command interface{}) (result interface{}, err error)

	// SendAndWaitWithTimeout Sends the given command and wait for it to execute. The result of the execution is returned when available.
	// This method will block until a result is available, or the given {@code timeout} was reached,or until the
	// thread is interrupted. When the timeout is reached or the thread is interrupted,this method returns nil
	// and the detail error.
	SendAndWaitWithTimeout(ctx context.Context, Command interface{}, timeout int64) (result interface{}, err error)
}
