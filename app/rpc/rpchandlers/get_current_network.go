package rpchandlers

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/Nirvana-Chain/nirvanad/app/rpc/rpccontext"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
