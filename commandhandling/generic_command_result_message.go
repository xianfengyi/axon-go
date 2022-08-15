package commandhandling

type GenericCommandResultMessage struct {
	resp interface{}
	err  error
}

func (g *GenericCommandResultMessage) GetResp() interface{} {
	return g.resp
}

func (g *GenericCommandResultMessage) GetError() error {
	return g.err
}

func (g *GenericCommandResultMessage) IsError() bool {
	if g.err == nil {
		return false
	}
	return true
}

// NewGenericCommandResultMessage creates a new result message.
func NewGenericCommandResultMessage(respDetail interface{}, detailErr error) CommandResultMessage {
	resultMessage := &GenericCommandResultMessage{
		resp: respDetail,
		err:  detailErr,
	}
	return resultMessage
}
