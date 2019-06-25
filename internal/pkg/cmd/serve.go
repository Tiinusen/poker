package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// GRPCServiceRegisterer ...
type GRPCServiceRegisterer func(*grpc.Server)

// ServeGRPC ...
func ServeGRPC(port int, grpcServiceRegisterer GRPCServiceRegisterer) error {

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(getBuiltinCert())),
	)
	reflection.Register(grpcServer)

	grpcServiceRegisterer(grpcServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()

	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrapf(err, "error listening to port %d", port)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*getBuiltinCert()},
			NextProtos:   []string{"h2"},
		},
	}

	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	if err != nil {
		return errors.Wrapf(err, "error serving on port %d", port)
	}

	return nil
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
