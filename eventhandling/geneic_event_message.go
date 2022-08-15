package eventhandling

import (
	"axon-go/messaging"
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"time"
)

// GenericEventMessage event is an internal representation of an event, returned when the aggregate
// uses NewEvent to create a new event. The events loaded from the db is
// represented by each DBs internal event type, implementing Event.
type GenericEventMessage struct {
	identifier    string
	eventName     string
	metaData      []byte
	payLoad       interface{}
	timestamp     time.Time
	aggregateName string
	aggregateId   string
	version       int
}

func (e GenericEventMessage) GetIdentifier() string {
	return ""
}

func (e GenericEventMessage) GetMetaData() []byte {
	return e.metaData
}

func (e GenericEventMessage) GetPayload() messaging.MessageData {
	return e.payLoad
}

func (e GenericEventMessage) GetPayloadType() reflect.Type {
	return reflect.TypeOf(e.payLoad)
}

func (e GenericEventMessage) GetEventName() string {
	return e.eventName
}

func (e GenericEventMessage) GetTimestamp() time.Time {
	return e.timestamp
}

func (e GenericEventMessage) GetAggregateName() string {
	return e.aggregateName
}

func (e GenericEventMessage) GetAggregateId() string {
	return e.aggregateId
}

func (e GenericEventMessage) GetVersion() int {
	return e.version
}

// NewGenericEventMessage creates a new command message from a business command.
func NewGenericEventMessage(eventName string, event interface{}) *GenericEventMessage {
	data, _ := json.Marshal(event)
	return &GenericEventMessage{
		identifier:    uuid.New().String(),
		eventName:     eventName,
		metaData:      data,
		payLoad:       event,
		timestamp:     time.Now(),
		aggregateName: "",
		aggregateId:   "",
		version:       0,
	}
}

// NewGenericEventMessageForAggregate creates a new event with a type and data, setting its
// timestamp. It also sets the aggregate data on it.
func NewGenericEventMessageForAggregate(eventName string, event interface{}, timestamp time.Time,
	aggregateName string, aggregateId string, version int) *GenericEventMessage {
	data, _ := json.Marshal(event)
	return &GenericEventMessage{
		identifier:    "",
		eventName:     eventName,
		metaData:      data,
		payLoad:       event,
		timestamp:     time.Now(),
		aggregateName: aggregateName,
		aggregateId:   aggregateId,
		version:       version,
	}
}
