package service

import (
	"log/slog"
	"math/rand"
	"skeleton-grpc/internal/lib"
	"sync"
)

type MockDataService struct {
	mx           sync.Mutex
	itemsCatalog map[int32]float32
	itemsAvail   map[int32]*lib.DataProduct
	log          *slog.Logger
}

func NewMockDataService(logger *slog.Logger) *MockDataService {
	itemCatalog := make(map[int32]float32)
	itemAvail := make(map[int32]*lib.DataProduct)

	return &MockDataService{
		itemsCatalog: itemCatalog,
		itemsAvail:   itemAvail,
		log:          logger,
	}
}

func (md *MockDataService) GetPriceService(id int32) (float32, error) {
	md.mx.Lock()
	defer md.mx.Unlock()
	md.itemsCatalog[id] = rand.Float32()

	return md.itemsCatalog[id], nil
}

func (md *MockDataService) GetAvailabilityService(id int32) (*lib.DataProduct, error) {
	md.mx.Lock()
	defer md.mx.Unlock()

	md.itemsAvail[id] = lib.NewDataProduct()

	md.log.Debug("GetAvailabilityService", md.itemsAvail)
	return md.itemsAvail[id], nil
}
