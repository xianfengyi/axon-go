package eventhandling

import (
	"context"
)

// EventBus sends published events to one of each handler type and all observers.
// That means that if the same handler is registered on multiple nodes only one
// of them will receive the event. In contrast all observers registered on multiple
// nodes will receive the event. GetEvents are not garantued to be handeled or observed
// in order.
type EventBus interface {
	// Publish the event on the bus.
	Publish(ctx context.Context, eventMessage GenericEventMessage) error

	// Subscribe the event on the bus
	Subscribe(eventName string, handler EventHandlerFunc) error
}
