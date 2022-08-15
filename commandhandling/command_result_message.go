package commandhandling

type CommandResultMessage interface {
	GetResp() interface{}
	GetError() error
	IsError() bool
}

type commandResultMessage struct {
	resp interface{}
	err  error
}

func (g commandResultMessage) GetResp() interface{} {
	return g.resp
}

func (g commandResultMessage) IsError() bool {
	if g.err == nil {
		return false
	}
	return true
}

func (g commandResultMessage) GetError() error {
	return g.err
}

// NewCommandResultMessage creates a new result message.
func NewCommandResultMessage(respDetail interface{}, detailErr error) CommandResultMessage {
	resultMessage := &commandResultMessage{
		resp: respDetail,
		err:  detailErr,
	}
	return resultMessage
}
