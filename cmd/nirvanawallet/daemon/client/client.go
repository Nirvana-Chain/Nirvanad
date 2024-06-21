package client

import (
	"context"
	"time"

	"github.com/Nirvana-Chain/nirvanad/cmd/nirvanawallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/Nirvana-Chain/nirvanad/cmd/nirvanawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the nirvanawalletd server, and returns the client instance
func Connect(address string) (pb.NirvanawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("nirvanawallet daemon is not running, start it with `nirvanawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewNirvanawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
