package modelling

import (
	"axon-go/config"
	"axon-go/eventhandling"
	"context"
	"errors"
	"reflect"
	"sync"
)

var aggregatesMu sync.RWMutex

// ErrAggregateNotRegistered is when no aggregate factory was registered.
var ErrAggregateNotRegistered = errors.New("aggregate not registered")

// ErrAggregateNotFound is when no aggregate can be found.
var ErrAggregateNotFound = errors.New("aggregate not found")

// The aggregate is created/loaded and saved by the Repository inside the
// Dispatcher. A domain specific aggregate can either implement the full interface,
// or more commonly embed *BaseAggregate to take care of the common methods.
type Aggregate interface {
	// GetAggregateId returns the id of the aggregate.
	GetAggregateId() string

	// GetAggregateType returns the type name of the aggregate.
	GetAggregateType() string

	// GetAggregateRoot returns the data of the aggregate.
	GetAggregateRoot() interface{}

	// ApplyEvent applies an event on the aggregate by setting its values.
	// If there are no errors the version should be incremented by calling
	// IncrementVersion.
	ApplyEvent(ctx context.Context, event interface{}) error
}

type BaseAggregate struct {
	AggregateId UinId
	Type        string
}

func (a BaseAggregate) GetAggregateId() string {
	return a.AggregateId.Identifier
}

func (a BaseAggregate) GetAggregateType() string {
	return a.Type
}

func (a BaseAggregate) GetAggregateRoot() interface{} {
	panic("implement me")
}

func (a BaseAggregate) ApplyEvent(ctx context.Context, event interface{}) error {
	eventName := reflect.TypeOf(event).Name()
	eventMessage := eventhandling.NewGenericEventMessage(eventName, event)

	// dispatch event to event bus
	config.GlobalEbus.Publish(ctx, *eventMessage)
	return nil
}
