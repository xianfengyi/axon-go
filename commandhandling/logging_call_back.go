package commandhandling

import (
	"fmt"
)

// CommandCallback implementation that simply logs the results of a command.
type LoggingCallback struct {
}

func (l LoggingCallback) OnResult(command CommandMessage, commandResult CommandResultMessage) {
	if commandResult.IsError() {
		fmt.Println("Command resulted in exception: ", command.GetCommandName(), commandResult.IsError())
	} else {
		fmt.Println("Command executed successfully: ", command.GetCommandName())
	}

}
