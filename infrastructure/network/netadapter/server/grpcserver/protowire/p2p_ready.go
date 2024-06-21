package protowire

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NirvanadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NirvanadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *NirvanadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
