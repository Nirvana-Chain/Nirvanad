package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Nirvana-Chain/nirvanad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.NirvanadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.NirvanadMessage_BanRequest{}),
	reflect.TypeOf(protowire.NirvanadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
