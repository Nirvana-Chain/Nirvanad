package testing

import (
	"testing"
	"time"

	"github.com/Nirvana-Chain/nirvanad/app/protocol/flows/v5/addressexchange"

	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	peerpkg "github.com/Nirvana-Chain/nirvanad/app/protocol/peer"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus"
	"github.com/Nirvana-Chain/nirvanad/domain/consensus/utils/testutils"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/addressmanager"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/netadapter/router"
)

type fakeReceiveAddressesContext struct{}

func (f fakeReceiveAddressesContext) AddressManager() *addressmanager.AddressManager {
	return nil
}

func TestReceiveAddressesErrors(t *testing.T) {
	testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
		incomingRoute := router.NewRoute("incoming")
		outgoingRoute := router.NewRoute("outgoing")
		peer := peerpkg.New(nil)
		errChan := make(chan error)
		go func() {
			errChan <- addressexchange.ReceiveAddresses(fakeReceiveAddressesContext{}, incomingRoute, outgoingRoute, peer)
		}()

		_, err := outgoingRoute.DequeueWithTimeout(time.Second)
		if err != nil {
			t.Fatalf("DequeueWithTimeout: %+v", err)
		}

		// Sending addressmanager.GetAddressesMax+1 addresses should trigger a ban
		err = incomingRoute.Enqueue(appmessage.NewMsgAddresses(make([]*appmessage.NetAddress,
			addressmanager.GetAddressesMax+1)))
		if err != nil {
			t.Fatalf("Enqueue: %+v", err)
		}

		select {
		case err := <-errChan:
			checkFlowError(t, err, true, true, "address count exceeded")
		case <-time.After(time.Second):
			t.Fatalf("timed out after %s", time.Second)
		}
	})
}
