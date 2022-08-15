package config

import (
	"axon-go/commandhandling"
	"axon-go/commandhandling/gateway"
	"axon-go/eventhandling"
)

var (
	GlobalCommandGateway gateway.CommandGateway
	GlobalEbus           eventhandling.EventBus
	GlobalCbus           commandhandling.CommandBus
)

type AxonConfig interface {
	GetCommandGateway() gateway.CommandGateway

	GetEventBus() eventhandling.EventBus

	GetCommandBus() commandhandling.CommandBus
}

type DefaultAxonConfig struct {
	CommandGateway gateway.CommandGateway
	Ebus           eventhandling.EventBus
	Cbus           commandhandling.CommandBus
}

func (d DefaultAxonConfig) GetCommandGateway() gateway.CommandGateway {
	return d.CommandGateway
}

func (d DefaultAxonConfig) GetEventBus() eventhandling.EventBus {
	return d.Ebus
}

func (d DefaultAxonConfig) GetCommandBus() commandhandling.CommandBus {
	return d.Cbus
}

func NewDefaultAxonConfig() *DefaultAxonConfig {
	axonCofig := new(DefaultAxonConfig)
	axonCofig.Cbus = commandhandling.NewSimpleCommandBus()
	axonCofig.Ebus = eventhandling.NewLocalEventBus()
	axonCofig.CommandGateway = gateway.NewSimpleCommandGatewayWithBus(axonCofig.Cbus)
	initGlobalConfig(axonCofig.Cbus, axonCofig.CommandGateway, axonCofig.Ebus)
	return axonCofig
}

func initGlobalConfig(cbus commandhandling.CommandBus, commandGateway gateway.CommandGateway, ebus eventhandling.EventBus) {
	GlobalCommandGateway = commandGateway
	GlobalEbus = ebus
	GlobalCbus = cbus
}
