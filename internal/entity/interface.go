package entity

type OrderRepositoryInterface interface {
	Save(order *Order) (*Order, error)
	Find(id string) (*Order, error)
	// GetTotal() (int, error)
}
