package app

import (
	"fmt"
	"log"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
	Shutdown(time.Duration)
}
type server struct {
	grpcServer   *grpc.Server
	grpcListener net.Listener
}

func NewServer(option *Option) Server {
	grpcServer := grpc.NewServer(
		unaryInterceptor(),
		grpc.ConnectionTimeout(time.Duration(option.ConnectTimeout)*time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: time.Duration(option.KeepAlive) * time.Second,
		}),
	)

	grpcListener, err := newListener(option.Port)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start grpc listener: %#v", err))
	}

	return &server{
		grpcServer:   grpcServer,
		grpcListener: grpcListener,
	}
}

func (s *server) Serve() error {
	errors := make(chan error)

	go func() {
		log.Printf("start grpc server")
		errors <- s.grpcServer.Serve(s.grpcListener)
	}()

	return <-errors
}

func (s *server) Shutdown(period time.Duration) {
	gracefulStopChan := make(chan bool, 1)

	go func() {
		log.Printf("grpcserver:  graceful shutdown is running")
		s.grpcServer.GracefulStop()
		gracefulStopChan <- true
	}()

	t := time.NewTimer(period)
	select {
	case <-gracefulStopChan:
		log.Printf("grpcserver: graceful shutdown completed before timing out")
	case <-t.C:
		log.Printf("graceful shutdown failed timeout, closing pending RPCs")
		s.grpcServer.Stop()
	}
}

// unaryInterceptor unary用のInterceptor
func unaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_validator.UnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	))
}

func newListener(serverAddr string) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%s", serverAddr))
}
