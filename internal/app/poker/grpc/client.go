package grpc

import proto "github.com/Tiinusen/poker/pkg/proto/poker"

// NewClient ...
func NewClient(client proto.PokerServiceClient) *Wrapper {
	return &Wrapper{
		client: client,
	}
}

// Wrapper ...
type Wrapper struct {
	client proto.PokerServiceClient
}
