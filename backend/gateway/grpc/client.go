package grpc

import (
	"github.com/ProjectSMAA/commons/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPCClient establishes a connection to the gRPC server
func gRPCClient() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// CallGrpcService is a helper function to avoid repetition in gRPC calls
func CallGrpcService[T any](clientFunc func(protos.UserServiceClient) (T, error)) (T, error) {
	var result T
	conn, err := gRPCClient()
	if err != nil {
		return result, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := protos.NewUserServiceClient(conn)
	return clientFunc(client)
}
