package commandhandling

import "axon-go/messaging"

// CommandMessage Represents a Message carrying a command as its payload. These messages carry an intention to change application
type CommandMessage interface {
	messaging.Message

	// GetCommandName Returns the name of the command to execute.
	GetCommandName() string
}
