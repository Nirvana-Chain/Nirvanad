package protowire

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NirvanadMessage_RequestPruningPointAndItsAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NirvanadMessage_RequestPruningPointAndItsAnticone is nil")
	}
	return &appmessage.MsgRequestPruningPointAndItsAnticone{}, nil
}

func (x *NirvanadMessage_RequestPruningPointAndItsAnticone) fromAppMessage(_ *appmessage.MsgRequestPruningPointAndItsAnticone) error {
	return nil
}
