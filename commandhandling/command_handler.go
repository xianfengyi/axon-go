package commandhandling

import "context"

// CommandHandler specialization for handlers of command.
type CommandHandler interface {
	// HandleCommand Handles the given {@code command}.
	HandleCommand(ctx context.Context, command interface{}) (resp interface{}, err error)
}

// CommandHandlerFunc is a function that can be used as a command handler.
type CommandHandlerFunc func(ctx context.Context, command interface{}) (resp interface{}, err error)

// HandleCommand  implements the HandleEvent method of the CommandHandler.
func (handlerFunc CommandHandlerFunc) HandleCommand(ctx context.Context, command interface{}) (resp interface{}, err error) {
	return handlerFunc(ctx, command)
}
