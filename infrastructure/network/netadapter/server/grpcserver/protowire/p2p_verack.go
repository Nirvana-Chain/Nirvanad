package protowire

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NirvanadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NirvanadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *NirvanadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
