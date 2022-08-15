package messaging

// Interface for a component that processes Messages.
type MessageHandler interface {
	// Handles the given {@code message}.
	Handle(message Message) error
}
