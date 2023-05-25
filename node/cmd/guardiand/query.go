package guardiand

import (
	"context"

	gossipv1 "github.com/certusone/wormhole/node/pkg/proto/gossip/v1"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/wormhole-foundation/wormhole/sdk/vaa"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// TODO: should this use a different standard of signing messages, like https://eips.ethereum.org/EIPS/eip-712
var queryRequestPrefix = []byte("query_request_00000000000000000000|")

func queryRequestDigest(b []byte) ethCommon.Hash {
	return ethCrypto.Keccak256Hash(append(queryRequestPrefix, b...))
}

var allowedRequestor = ethCommon.BytesToAddress(ethCommon.Hex2Bytes("beFA429d57cD18b7F8A4d91A2da9AB4AF05d0FBe"))

// Multiplex observation requests to the appropriate chain
func handleQueryRequests(
	ctx context.Context,
	logger *zap.Logger,
	signedQueryReqC <-chan *gossipv1.SignedQueryRequest,
	chainQueryReqC map[vaa.ChainID]chan *gossipv1.SignedQueryRequest,
	enableFlag bool,
) {
	qLogger := logger.With(zap.String("component", "ccqhandler"))
	if enableFlag {
		qLogger.Info("cross chain queries are enabled")
	}

	for {
		select {
		case <-ctx.Done():
			return
		case signedQueryRequest := <-signedQueryReqC:
			if !enableFlag {
				qLogger.Error("received a query request when the feature is disabled, dropping it")
				continue
			}
			// requestor validation happens here
			// request type validation is currently handled by the watcher
			// in the future, it may be worthwhile to catch certain types of
			// invalid requests here for tracking purposes
			// e.g.
			// - length check on "signature" 65 bytes
			// - length check on "to" address 20 bytes
			// - valid "block" strings

			digest := queryRequestDigest(signedQueryRequest.QueryRequest)

			signerBytes, err := ethCrypto.Ecrecover(digest.Bytes(), signedQueryRequest.Signature)
			if err != nil {
				qLogger.Error("failed to recover public key")
				continue
			}

			signerAddress := ethCommon.BytesToAddress(ethCrypto.Keccak256(signerBytes[1:])[12:])

			if signerAddress != allowedRequestor {
				qLogger.Error("invalid requestor", zap.String("requestor", signerAddress.Hex()))
				continue
			}

			var queryRequest gossipv1.QueryRequest
			err = proto.Unmarshal(signedQueryRequest.QueryRequest, &queryRequest)
			if err != nil {
				qLogger.Error("received invalid message",
					zap.String("requestor", signerAddress.Hex()))
				continue
			}

			if channel, ok := chainQueryReqC[vaa.ChainID(queryRequest.ChainId)]; ok {
				select {
				// TODO: only send the query request itself and reassemble in this module
				case channel <- signedQueryRequest:
				default:
					qLogger.Warn("failed to send query request to watcher",
						zap.Uint16("chain_id", uint16(queryRequest.ChainId)))
				}
			} else {
				qLogger.Error("unknown chain ID for query request",
					zap.Uint16("chain_id", uint16(queryRequest.ChainId)))
			}
		}
	}
}
