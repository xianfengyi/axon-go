package commandhandling

import (
	"axon-go/messaging"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// ---------------------------------------------------
// --------mock data start----------------------------

type TestCommandMessage struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t TestCommandMessage) GetIdentifier() string {
	return "123"
}

func (t TestCommandMessage) GetMetaData() []byte {
	data, _ := json.Marshal(&t)
	return data
}

func (t TestCommandMessage) GetPayload() messaging.MessageData {
	return t
}

func (t TestCommandMessage) GetPayloadType() reflect.Type {
	return reflect.TypeOf(t)
}

func (t TestCommandMessage) GetCommandName() string {
	return "TestCommandMessage"
}

// --------mock data end----------------------------
// ---------------------------------------------------

func TestCreateCommandMessage(t *testing.T) {
	commandMessage := TestCommandMessage{
		Name: "pioneeryi",
		Age:  25,
	}
	assert.Equal(t, "TestCommandMessage", commandMessage.GetCommandName())
	assert.Equal(t, "123", commandMessage.GetIdentifier())
	assert.Equal(t, reflect.TypeOf(commandMessage), commandMessage.GetPayloadType())
	assert.NotEmpty(t, commandMessage.GetPayload())
}
