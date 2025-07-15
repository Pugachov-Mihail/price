package api

import (
	"google.golang.org/grpc"
	price "skeleton-grpc/internal/protos/gen/gen/protos"
)

type Price struct {
	price.UnimplementedPriceServiceServer
}

func RegisterPriceServer(registrar grpc.ServiceRegistrar) {
	price.RegisterPriceServiceServer(registrar, &Price{})
}
