package cmd

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

// DialGRPC ...
func DialGRPC(host string, timeout time.Duration) (*grpc.ClientConn, error) {
	if len(host) == 0 {
		return nil, fmt.Errorf("no hosts provided")
	}

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(getTrustedCerts(), "")),
	}

	conn, err := grpc.Dial(host, dialOpts...)
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()
	err = awaitConnect(ctx, conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func awaitConnect(ctx context.Context, conn *grpc.ClientConn) error {
	for {
		currState := conn.GetState()
		switch currState {
		case connectivity.Ready:
			return nil
		//case connectivity.TransientFailure: // No good because connection will enter transient failure state when attempting to connect to a server that is not yet up
		//	return errors.Errorf("grpc connection failure")
		// TODO figure out how to detect certificate errors, because right now it will just "time out" when server is up and running but with bad cert
		default:
			ok := conn.WaitForStateChange(ctx, currState)
			if !ok {
				return ctx.Err()
			}
		}
	}
}
