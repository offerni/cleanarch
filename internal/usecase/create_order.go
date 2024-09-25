package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type CreateOrderInputDTO struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	*FetchOrderOutputDTO // to force consistency among single-entity responses
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (CreateOrderOutputDTO, error) {
	order := entity.Order{
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	resp, err := c.OrderRepository.Save(&order)
	if err != nil {
		return CreateOrderOutputDTO{}, err
	}

	dto := CreateOrderOutputDTO{
		&FetchOrderOutputDTO{
			ID:         resp.ID,
			Price:      resp.Price,
			Tax:        resp.Tax,
			FinalPrice: resp.Price + resp.Tax,
		},
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
