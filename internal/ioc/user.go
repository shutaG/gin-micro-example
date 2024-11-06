package ioc

import (
	userv1 "gin-micro-example/api/proto/gen/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserClient() userv1.UserServiceClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc, err := grpc.NewClient("127.0.0.1:8090", opts...)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserServiceClient(cc)
}
