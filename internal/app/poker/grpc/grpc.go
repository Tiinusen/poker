package grpc

import (
	"context"

	"github.com/Tiinusen/poker/internal/app/poker"
	proto "github.com/Tiinusen/poker/pkg/proto/poker"
	"google.golang.org/grpc"
)

// NewServer ...
func NewServer(service *poker.Service) *Server {
	return &Server{
		service: service,
	}
}

// Server ...
type Server struct {
	service *poker.Service
}

var _ proto.PokerServiceClient = &Server{}

// Search ...
func (s *Server) Connect(ctx context.Context, opts ...grpc.CallOption) (proto.PokerService_ConnectClient, error) {
	// ctx := stream.Context()
	// personChan, err := s.service.Search(ctx, domain.ProtoSearchCriteriaToDomainSearchCriteriaN(req.Criterias))
	// if err != nil {
	// 	return err
	// }
	// for {
	// 	// select {
	// 	// case <-ctx.Done():
	// 	// 	return nil
	// 	// case person := <-personChan:
	// 	// 	stream.Send(domain.DomainPersonToProtoPerson(person))
	// 	// }
	// }
	return nil, nil
}
