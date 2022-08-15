package commandhandling

import (
	"axon-go/messaging"
	"axon-go/utils"
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
)

// GenericCommandMessage implementation of a {@link Message}, which convert a biz command to {@code CommandMessage}
type GenericCommandMessage struct {
	Identifier string
	Name       string
	MetaData   []byte
	PayLoad    interface{}
}

func (g *GenericCommandMessage) GetIdentifier() string {
	return g.Identifier
}

func (g *GenericCommandMessage) GetMetaData() []byte {
	return g.MetaData
}

func (g *GenericCommandMessage) GetPayload() messaging.MessageData {
	return g.PayLoad
}

func (g *GenericCommandMessage) GetPayloadType() reflect.Type {
	return reflect.TypeOf(g.PayLoad)
}

func (g *GenericCommandMessage) GetCommandName() string {
	return g.Name
}

// NewGenericCommandMessage creates a new command message from a business command.
func NewGenericCommandMessage(command interface{}) CommandMessage {
	data, _ := json.Marshal(command)

	// Because struct and interface get name is different, so we first try get name, if nil, then get elem name
	commandName := utils.GetObjectName(command)
	// construct GenericCommandMessage
	commandMessage := &GenericCommandMessage{
		Identifier: uuid.New().String(),
		Name:       commandName,
		MetaData:   data,
		PayLoad:    command,
	}
	return commandMessage
}
