package entity

type OrderRepositoryInterface interface {
	Save(order *Order) (*Order, error)
	Find(id string) (*Order, error)
	FindAll() ([]*Order, error) // in a real world scenarion we'd have an account_id here
}
