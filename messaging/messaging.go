package messaging

import "reflect"

// EventData is any additional data for an event.
type MessageData interface{}

// Representation of a Message, containing a Payload. Typical examples of Messages are Commands and GetEvents.
type Message interface {
	// Returns the identifier of this message.
	GetIdentifier() string

	// Returns the meta data for this message.
	GetMetaData() []byte

	// Returns the payload of this message.
	GetPayload() MessageData

	// Returns the type of the payload.
	GetPayloadType() reflect.Type
}
