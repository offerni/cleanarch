package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type FetchOrderInputDTO struct {
	ID string `json:"id"`
}

type FetchOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type FetchOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFetchOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *FetchOrderUseCase {
	return &FetchOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (f *FetchOrderUseCase) Execute(input FetchOrderInputDTO) (FetchOrderOutputDTO, error) {
	order, err := f.OrderRepository.Find(input.ID)
	if err != nil {
		return FetchOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}, err
	}

	dto := FetchOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	return dto, nil
}
