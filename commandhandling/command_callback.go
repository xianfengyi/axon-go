package commandhandling

// CommandCallback Interface describing a callback that is invoked when command handler execution has finished.
type CommandCallback interface {
	// OnResult Invoked when command handling execution is completed.
	OnResult(command CommandMessage, commandResult CommandResultMessage)
}

type FutureCallback struct {
	resp interface{}
	err  error
}

func (f *FutureCallback) OnResult(command CommandMessage, commandResult CommandResultMessage) {
	f.resp = commandResult.GetResp()
	f.err = commandResult.GetError()
}

func (f *FutureCallback) GetResult() (resp interface{}, err error) {
	return f.resp, f.err
}

func NewFutureCallback() *FutureCallback {
	return new(FutureCallback)
}
