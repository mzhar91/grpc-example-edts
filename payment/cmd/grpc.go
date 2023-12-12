package cmd

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	_config "github.com/grpc-example-edts/payment/config"
	_load "github.com/grpc-example-edts/payment/config/load"
)

func StartServerGRPC(connection *_config.Connection, timeoutContext time.Duration) {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf("%v%v", _config.Env.ServerHostGRPC, _config.Env.ServerAddrGRPC))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))

	grpcServer := grpc.NewServer(opts...)
	_load.GrpcLoad(grpcServer, connection, timeoutContext)

	fmt.Println(fmt.Sprintf("⇨ grpc server started on %s\n", _config.Env.ServerAddrGRPC))

	grpcServer.Serve(listen)

}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)

	// claimsRpc := *_authHelper.InitClaimsRpc(_config.Env.Debug)
	// newCtx, err, _ := claimsRpc.ClaimsRpc(ctx)
	// if err != nil {
	// 	return nil, status.Errorf(
	// 		codes.Canceled,
	// 		fmt.Sprintf("%v", err),
	// 	)
	// }

	return handler(ctx, req)
}
