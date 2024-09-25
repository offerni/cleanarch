package database

import (
	"database/sql"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/google/uuid"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) (*entity.Order, error) {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return nil, err
	}

	return &entity.Order{
		ID:         id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

func (r *OrderRepository) Find(id string) (*entity.Order, error) {
	order := entity.Order{}
	err := r.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", id).Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
