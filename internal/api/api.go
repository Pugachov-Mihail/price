package api

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"skeleton-grpc/internal/lib"
	price "skeleton-grpc/internal/protos/gen/gen/protos"
)

type Price struct {
	PriceServ ServicePrice
	price.UnimplementedPriceServiceServer
}

type ServicePrice interface {
	GetPriceService(id int32) (float32, error)
	GetAvailabilityService(id int32) (*lib.DataProduct, error)
}

func RegisterPriceServer(registrar grpc.ServiceRegistrar, srv ServicePrice) {
	price.RegisterPriceServiceServer(registrar, &Price{PriceServ: srv})
}

func (p *Price) GetPrice(ctx context.Context, req *price.ItemsPriceRequest) (*price.ItemsPriceResponce, error) {
	if req.GetItemId() == 0 {
		return &price.ItemsPriceResponce{}, status.Error(codes.InvalidArgument, "Item ID should be zero")
	}

	res, err := p.PriceServ.GetPriceService(req.GetItemId())
	if err != nil {
		return &price.ItemsPriceResponce{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &price.ItemsPriceResponce{
		ItemId: req.GetItemId(),
		Price:  res,
	}, nil
}

func (p *Price) GetAvailability(ctx context.Context, req *price.ItemsPriceRequest) (*price.ItemsAvailabilityResponce, error) {
	if req.GetItemId() == 0 {
		return &price.ItemsAvailabilityResponce{}, status.Error(codes.InvalidArgument, "Item ID should be zero")
	}

	res, err := p.PriceServ.GetAvailabilityService(req.GetItemId())
	if err != nil {
		return &price.ItemsAvailabilityResponce{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &price.ItemsAvailabilityResponce{ItemId: req.GetItemId(), Cdate: res.Cdate, Mdate: res.Mdate, Count: res.Count}, nil
}
