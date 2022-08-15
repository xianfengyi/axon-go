package commandhandling

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// ---------------------------------------------------
// --------mock data start----------------------------

type TestCommandResultMessage struct {
	Result string `json:"result"`
}

func (t TestCommandResultMessage) GetIdentifier() string {
	return "123"
}

func (t TestCommandResultMessage) GetPayload() []byte {
	data, _ := json.Marshal(&t)
	return data
}

func (t TestCommandResultMessage) GetPayloadType() reflect.Type {
	return reflect.TypeOf(t)
}

// --------mock data end----------------------------
// ---------------------------------------------------

func TestCreateCommandResultMessage(t *testing.T) {
	commandMessage := TestCommandResultMessage{
		Result: "success",
	}
	assert.Equal(t, "123", commandMessage.GetIdentifier())
	assert.Equal(t, reflect.TypeOf(commandMessage), commandMessage.GetPayloadType())
	assert.NotEmpty(t, commandMessage.GetPayload())
}
