package eventhandling

import (
	"context"
)

// EventHandler is a handler of events. If registered on a bus as a handler only
// one handler of the same type will receive each event. If registered on a bus
// as an observer all handlers of the same type will receive each event.
type EventHandler interface {
	// HandleEvent handles an event.
	HandleEvent(ctx context.Context, eventMessage EventMessage)
}

// EventHandlerFunc is a function that can be used as a event handler.
type EventHandlerFunc func(ctx context.Context, eventMessage EventMessage)

// HandleEvent implements the HandleEvent method of the EventHandler.
func (h EventHandlerFunc) HandleEvent(ctx context.Context, eventMessage EventMessage) {
	h(ctx, eventMessage)
}
