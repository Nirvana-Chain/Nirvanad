package protowire

import (
	"github.com/Nirvana-Chain/nirvanad/app/appmessage"
	"github.com/pkg/errors"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

// ToAppMessage converts a NirvanadMessage to its appmessage.Message representation
func (x *NirvanadMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NirvanadMessage is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

// FromAppMessage creates a NirvanadMessage from a appmessage.Message
func FromAppMessage(message appmessage.Message) (*NirvanadMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &NirvanadMessage{
		Payload: payload,
	}, nil
}

func toPayload(message appmessage.Message) (isNirvanadMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isNirvanadMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.MsgAddresses:
		payload := new(NirvanadMessage_Addresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlock:
		payload := new(NirvanadMessage_Block)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestBlockLocator:
		payload := new(NirvanadMessage_RequestBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockLocator:
		payload := new(NirvanadMessage_BlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAddresses:
		payload := new(NirvanadMessage_RequestAddresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDBlocks:
		payload := new(NirvanadMessage_RequestIBDBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextHeaders:
		payload := new(NirvanadMessage_RequestNextHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneHeaders:
		payload := new(NirvanadMessage_DoneHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestRelayBlocks:
		payload := new(NirvanadMessage_RequestRelayBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestTransactions:
		payload := new(NirvanadMessage_RequestTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTransactionNotFound:
		payload := new(NirvanadMessage_TransactionNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvRelayBlock:
		payload := new(NirvanadMessage_InvRelayBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvTransaction:
		payload := new(NirvanadMessage_InvTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPing:
		payload := new(NirvanadMessage_Ping)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPong:
		payload := new(NirvanadMessage_Pong)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTx:
		payload := new(NirvanadMessage_Transaction)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVerAck:
		payload := new(NirvanadMessage_Verack)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVersion:
		payload := new(NirvanadMessage_Version)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReject:
		payload := new(NirvanadMessage_Reject)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointUTXOSet:
		payload := new(NirvanadMessage_RequestPruningPointUTXOSet)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointUTXOSetChunk:
		payload := new(NirvanadMessage_PruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgUnexpectedPruningPoint:
		payload := new(NirvanadMessage_UnexpectedPruningPoint)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocator:
		payload := new(NirvanadMessage_IbdBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHash:
		payload := new(NirvanadMessage_IbdBlockLocatorHighestHash)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHashNotFound:
		payload := new(NirvanadMessage_IbdBlockLocatorHighestHashNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockHeadersMessage:
		payload := new(NirvanadMessage_BlockHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointUTXOSetChunk:
		payload := new(NirvanadMessage_RequestNextPruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDonePruningPointUTXOSetChunks:
		payload := new(NirvanadMessage_DonePruningPointUtxoSetChunks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedData:
		payload := new(NirvanadMessage_BlockWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointAndItsAnticone:
		payload := new(NirvanadMessage_RequestPruningPointAndItsAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneBlocksWithTrustedData:
		payload := new(NirvanadMessage_DoneBlocksWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlock:
		payload := new(NirvanadMessage_IbdBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestHeaders:
		payload := new(NirvanadMessage_RequestHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPoints:
		payload := new(NirvanadMessage_PruningPoints)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointProof:
		payload := new(NirvanadMessage_RequestPruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointProof:
		payload := new(NirvanadMessage_PruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReady:
		payload := new(NirvanadMessage_Ready)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTrustedData:
		payload := new(NirvanadMessage_TrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedDataV4:
		payload := new(NirvanadMessage_BlockWithTrustedDataV4)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks:
		payload := new(NirvanadMessage_RequestNextPruningPointAndItsAnticoneBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDChainBlockLocator:
		payload := new(NirvanadMessage_RequestIBDChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDChainBlockLocator:
		payload := new(NirvanadMessage_IbdChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAnticone:
		payload := new(NirvanadMessage_RequestAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}

func toRPCPayload(message appmessage.Message) (isNirvanadMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.GetCurrentNetworkRequestMessage:
		payload := new(NirvanadMessage_GetCurrentNetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCurrentNetworkResponseMessage:
		payload := new(NirvanadMessage_GetCurrentNetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockRequestMessage:
		payload := new(NirvanadMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(NirvanadMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(NirvanadMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(NirvanadMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedRequestMessage:
		payload := new(NirvanadMessage_NotifyBlockAddedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedResponseMessage:
		payload := new(NirvanadMessage_NotifyBlockAddedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockAddedNotificationMessage:
		payload := new(NirvanadMessage_BlockAddedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesRequestMessage:
		payload := new(NirvanadMessage_GetPeerAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesResponseMessage:
		payload := new(NirvanadMessage_GetPeerAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashRequestMessage:
		payload := new(NirvanadMessage_GetSelectedTipHashRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashResponseMessage:
		payload := new(NirvanadMessage_GetSelectedTipHashResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryRequestMessage:
		payload := new(NirvanadMessage_GetMempoolEntryRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryResponseMessage:
		payload := new(NirvanadMessage_GetMempoolEntryResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoRequestMessage:
		payload := new(NirvanadMessage_GetConnectedPeerInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoResponseMessage:
		payload := new(NirvanadMessage_GetConnectedPeerInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerRequestMessage:
		payload := new(NirvanadMessage_AddPeerRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerResponseMessage:
		payload := new(NirvanadMessage_AddPeerResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionRequestMessage:
		payload := new(NirvanadMessage_SubmitTransactionRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionResponseMessage:
		payload := new(NirvanadMessage_SubmitTransactionResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage:
		payload := new(NirvanadMessage_NotifyVirtualSelectedParentChainChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage:
		payload := new(NirvanadMessage_NotifyVirtualSelectedParentChainChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentChainChangedNotificationMessage:
		payload := new(NirvanadMessage_VirtualSelectedParentChainChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockRequestMessage:
		payload := new(NirvanadMessage_GetBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockResponseMessage:
		payload := new(NirvanadMessage_GetBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkRequestMessage:
		payload := new(NirvanadMessage_GetSubnetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkResponseMessage:
		payload := new(NirvanadMessage_GetSubnetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage:
		payload := new(NirvanadMessage_GetVirtualSelectedParentChainFromBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage:
		payload := new(NirvanadMessage_GetVirtualSelectedParentChainFromBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksRequestMessage:
		payload := new(NirvanadMessage_GetBlocksRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksResponseMessage:
		payload := new(NirvanadMessage_GetBlocksResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountRequestMessage:
		payload := new(NirvanadMessage_GetBlockCountRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountResponseMessage:
		payload := new(NirvanadMessage_GetBlockCountResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(NirvanadMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(NirvanadMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictRequestMessage:
		payload := new(NirvanadMessage_ResolveFinalityConflictRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictResponseMessage:
		payload := new(NirvanadMessage_ResolveFinalityConflictResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsRequestMessage:
		payload := new(NirvanadMessage_NotifyFinalityConflictsRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsResponseMessage:
		payload := new(NirvanadMessage_NotifyFinalityConflictsResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictNotificationMessage:
		payload := new(NirvanadMessage_FinalityConflictNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictResolvedNotificationMessage:
		payload := new(NirvanadMessage_FinalityConflictResolvedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesRequestMessage:
		payload := new(NirvanadMessage_GetMempoolEntriesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesResponseMessage:
		payload := new(NirvanadMessage_GetMempoolEntriesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownRequestMessage:
		payload := new(NirvanadMessage_ShutDownRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownResponseMessage:
		payload := new(NirvanadMessage_ShutDownResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersRequestMessage:
		payload := new(NirvanadMessage_GetHeadersRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersResponseMessage:
		payload := new(NirvanadMessage_GetHeadersResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedRequestMessage:
		payload := new(NirvanadMessage_NotifyUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedResponseMessage:
		payload := new(NirvanadMessage_NotifyUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UTXOsChangedNotificationMessage:
		payload := new(NirvanadMessage_UtxosChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedRequestMessage:
		payload := new(NirvanadMessage_StopNotifyingUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedResponseMessage:
		payload := new(NirvanadMessage_StopNotifyingUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesRequestMessage:
		payload := new(NirvanadMessage_GetUtxosByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesResponseMessage:
		payload := new(NirvanadMessage_GetUtxosByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(NirvanadMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(NirvanadMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage:
		payload := new(NirvanadMessage_GetVirtualSelectedParentBlueScoreRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage:
		payload := new(NirvanadMessage_GetVirtualSelectedParentBlueScoreResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage:
		payload := new(NirvanadMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage:
		payload := new(NirvanadMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage:
		payload := new(NirvanadMessage_VirtualSelectedParentBlueScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanRequestMessage:
		payload := new(NirvanadMessage_BanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanResponseMessage:
		payload := new(NirvanadMessage_BanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanRequestMessage:
		payload := new(NirvanadMessage_UnbanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanResponseMessage:
		payload := new(NirvanadMessage_UnbanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoRequestMessage:
		payload := new(NirvanadMessage_GetInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoResponseMessage:
		payload := new(NirvanadMessage_GetInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage:
		payload := new(NirvanadMessage_NotifyPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage:
		payload := new(NirvanadMessage_NotifyPruningPointUTXOSetOverrideResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.PruningPointUTXOSetOverrideNotificationMessage:
		payload := new(NirvanadMessage_PruningPointUTXOSetOverrideNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage:
		payload := new(NirvanadMessage_StopNotifyingPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(NirvanadMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(NirvanadMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedRequestMessage:
		payload := new(NirvanadMessage_NotifyVirtualDaaScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedResponseMessage:
		payload := new(NirvanadMessage_NotifyVirtualDaaScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualDaaScoreChangedNotificationMessage:
		payload := new(NirvanadMessage_VirtualDaaScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesRequestMessage:
		payload := new(NirvanadMessage_GetBalancesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesResponseMessage:
		payload := new(NirvanadMessage_GetBalancesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(NirvanadMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(NirvanadMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(NirvanadMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesRequestMessage:
		payload := new(NirvanadMessage_GetMempoolEntriesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesResponseMessage:
		payload := new(NirvanadMessage_GetMempoolEntriesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyRequestMessage:
		payload := new(NirvanadMessage_GetCoinSupplyRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyResponseMessage:
		payload := new(NirvanadMessage_GetCoinSupplyResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}
