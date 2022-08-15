package eventhandling

import (
	"axon-go/messaging"
	"time"
)

// EventMessage Event is a domain event describing a change that has happened to an aggregate.
// The event should contain all the data needed when applying/handling it.
type EventMessage interface {
	// Implements messaging.Message
	messaging.Message

	// GetEventName returns the name of the event.
	GetEventName() string

	// Timestamp of when the event was created.
	GetTimestamp() time.Time

	// AggregateName returns the name of the aggregate that the event can be applied to.
	GetAggregateName() string

	// GetAggregateId returns the id of the aggregate that the event should be applied to.
	GetAggregateId() string

	// GetVersion of the aggregate for this event (after it has been applied).
	GetVersion() int
}

// NewEventMessage creates a new event with a type and data, setting its timestamp.
func NewEventMessage(eventName string, event interface{}) EventMessage {
	return NewGenericEventMessage(eventName, event)
}

// NewEventForAggregate creates a new event with a type and data, setting its
// timestamp. It also sets the aggregate data on it.
func NewEventForAggregate(eventName string, event interface{}, timestamp time.Time,
	aggregateName string, aggregateId string, version int) EventMessage {
	return NewGenericEventMessageForAggregate(eventName, event, timestamp, aggregateName, aggregateId, version)
}
