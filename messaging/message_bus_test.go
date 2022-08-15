package messaging

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var testmessageBus *messageBus

func init() {
	testmessageBus = new(messageBus)
	testmessageBus.handlerQueueSize = 1000
	testmessageBus.handlers = make(map[string][]*handler)
}

func TestPublish(t *testing.T) {

	var topic = "test"

	myHandler := handler{
		queue: make(chan []reflect.Value),
	}

	var handlers []*handler
	handlers = append(handlers, &myHandler)
	testmessageBus.handlers[topic] = handlers

	go func() {
		out := <-handlers[0].queue
		fmt.Println(out)
	}()

	testmessageBus.Publish(topic, 1, 2, 3)
}

func TestBuildHandlerArgs(t *testing.T) {
	var args []interface{}
	args = append(args, 1, 2, 3)

	result := buildHandlerArgs(args)
	assert.Equal(t, 3, len(result))
}

func TestSubscribe(t *testing.T) {
	var topic = "test"

	err := testmessageBus.Subscribe(topic, mockHandleFunc)
	assert.Nil(t, err)
}

func mockHandleFunc(messageType string) {
	fmt.Println("it is mockHandleFunc")
}
