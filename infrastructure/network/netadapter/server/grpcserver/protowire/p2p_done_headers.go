package protowire

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NirvanadMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NirvanadMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *NirvanadMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
