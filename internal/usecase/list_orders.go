package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersOutputDTO struct {
	Data []*FetchOrderOutputDTO
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (f *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	res, err := f.OrderRepository.FindAll()

	orders := make([]*FetchOrderOutputDTO, len(res))

	for i, order := range res {
		orders[i] = &FetchOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	if err != nil {
		return ListOrdersOutputDTO{
			Data: []*FetchOrderOutputDTO{},
		}, err
	}

	dto := ListOrdersOutputDTO{
		Data: orders,
	}

	return dto, nil
}
