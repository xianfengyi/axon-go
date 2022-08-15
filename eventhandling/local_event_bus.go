package eventhandling

import (
	"axon-go/messaging"
	"context"
	"runtime"
)

// LocalEventBus a local event bus that delegates handling of published events
// to all matching registered handlers, in order of registration.
type LocalEventBus struct {
	messageBus messaging.MessageBus
}

// NewLocalEventBus creates a EventBus.
func NewLocalEventBus() EventBus {
	return &LocalEventBus{
		messageBus: messaging.New(runtime.NumCPU()),
	}
}

// Publish implements the method of the bus.EventBus interface.
func (bus *LocalEventBus) Publish(ctx context.Context, eventMessage GenericEventMessage) error {
	bus.messageBus.Publish(eventMessage.GetEventName(), ctx, eventMessage)
	return nil
}

// Subscribe implements the method of the bus.EventBus interface.
func (bus *LocalEventBus) Subscribe(eventName string, handler EventHandlerFunc) error {
	return bus.messageBus.Subscribe(eventName, handler)
}
