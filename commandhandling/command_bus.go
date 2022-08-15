package commandhandling

import "context"

// CommandBus allows to subscribe/dispatch commands
type CommandBus interface {
	// Dispatch the given {@code CommandMessage} to the CommandHandler subscribed to the given {@code command}'s name.
	// No feedback is given about the status of the dispatching process. Implementations may return immediately after
	// asserting a valid handler is registered for the given command.
	Dispatch(ctx context.Context, command CommandMessage) error

	// DispatchWithCallback Dispatch the given {@code CommandMessage} to the CommandHandler subscribed to the given {@code CommandMessage}'s name.
	// When the command is processed, one of the callback's methods is called, depending on the result of the processing.
	DispatchWithCallback(ctx context.Context, command CommandMessage, callback CommandCallback) error

	// Subscribe the given {@code CommandHandler} to commands with the given {@code commandName}.
	Subscribe(commandName string, handler CommandHandler) error
}
