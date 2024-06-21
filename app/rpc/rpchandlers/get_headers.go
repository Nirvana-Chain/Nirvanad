package rpchandlers

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/Nirvana-Chain/nirvanad/app/rpc/rpccontext"
	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
