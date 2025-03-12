package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func gRPCClient(host, port string) (*grpc.ClientConn, error) {
	return grpc.NewClient(host+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func CallGrpcService[C any, T any](host, port string,
	clientCreator func(conn grpc.ClientConnInterface) C,
	call func(C) (T, error)) (T, error) {
	var result T
	conn, err := gRPCClient(host, port)
	if err != nil {
		return result, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := clientCreator(conn)
	return call(client)
}
