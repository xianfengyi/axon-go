# axon-go
A simple axon framework for GO language implementation

## Features
* domain model
* command bus
* command gateway
* event bus

## Usage
### Config
In general, we use default configuration. The default command bus is synchronous
,and the default event bus is local event bus which implemented by chan of golang.

In you code ,you just need:
```go
var AxonConfig config.AxonConfig

func init() {
	AxonConfig = config.NewDefaultAxonConfig()
}
```
### Pub/Sub Command
#### Publish Command
The param domainCommand is the service customized command, and AxonConfig is the default axon configuration. 
Code like this:
```go
resp, err := AxonConfig.GetCommandGateway().SendAndWait(contxt, domainCommand)
```
#### Subscribe Command
The framework uses the command name as topic by default, so the acceptance also needs to use the command name. 
In order to avoid writing the name wrong, it is recommended to use reflection to obtain the name.
The subscriber uses a custom CommandHandler.
```go
_ = AxonConfig.GetCommandBus().Subscribe(reflect.TypeOf(customCommand).Name(), CustomCommandHandler())
```
The Command Handler interface is defined as follows:
```go
type CommandHandlerFunc func(ctx context.Context, command interface{}) (resp interface{}, err error)
```
A custom CommandHandler has the following general form:
```go
func CustomCommandHandler() commandhandling.CommandHandlerFunc {
    return func(ctx context.Context, commandMessage interface{}) (result interface{}, err error) {
    
    // 1） get command data
    
    // 2） operate aggregate	
    
    return resp, nil
    }
}
```

### Pub/Sub Event
#### Apply Event
The param customEvent is the event defined by the service. The events published using the current aggregation object are as follows:
```go
aggregate.ApplyEvent(context, customEvent)
```
#### Subscribe
The framework uses the event name as topic by default, so the subscription also needs to use the event name. 
In order to avoid writing the name wrong, it is suggested to use reflection to obtain the name, and the subscriber is a custom EventHandler.
```go
_ = AxonConfig.GetEventBus().Subscribe(reflect.TypeOf(customeEvent).Name(), customeEventHandler)
```
the event handler is like this:
```go
func CustomeEventHandler() eventhandling.EventHandlerFunc {
	return func(ctx context.Context, eventMessage eventhandling.EventMessage) { 
		// 1） get event data
        
        // 2） do your business	
	}
}
```

### Create Aggregate
```go
type CustomAggregate struct {
	  modelling.BaseAggregate
	  
	  // define aggregate member
}
```

## TODO
* more event bus implementation
* event store
* event snapshot