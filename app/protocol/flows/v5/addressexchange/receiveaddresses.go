package addressexchange

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/Nirvana-Chain/nirvanad/app/protocol/common"
	peerpkg "github.com/Nirvana-Chain/nirvanad/app/protocol/peer"
	"github.com/Nirvana-Chain/nirvanad/app/protocol/protocolerrors"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/addressmanager"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/netadapter/router"
)

// ReceiveAddressesContext is the interface for the context needed for the ReceiveAddresses flow.
type ReceiveAddressesContext interface {
	AddressManager() *addressmanager.AddressManager
}

// ReceiveAddresses asks a peer for more addresses if needed.
func ReceiveAddresses(context ReceiveAddressesContext, incomingRoute *router.Route, outgoingRoute *router.Route,
	peer *peerpkg.Peer) error {

	subnetworkID := peer.SubnetworkID()
	msgGetAddresses := appmessage.NewMsgRequestAddresses(false, subnetworkID)
	err := outgoingRoute.Enqueue(msgGetAddresses)
	if err != nil {
		return err
	}

	message, err := incomingRoute.DequeueWithTimeout(common.DefaultTimeout)
	if err != nil {
		return err
	}

	msgAddresses := message.(*appmessage.MsgAddresses)
	if len(msgAddresses.AddressList) > addressmanager.GetAddressesMax {
		return protocolerrors.Errorf(true, "address count exceeded %d", addressmanager.GetAddressesMax)
	}

	return context.AddressManager().AddAddresses(msgAddresses.AddressList...)
}
