package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgOrderRepository struct {
	pool *pgxpool.Pool
}

func NewPgOrderRepository(pool *pgxpool.Pool) *PgOrderRepository {
	return &PgOrderRepository{pool: pool}
}

func (r *PgOrderRepository) Create(ctx context.Context, order entity.Order) (entity.Order, error) {
	const query = `
		INSERT INTO orders
		(status, client_id, item_id, item_count, total_price)
		values ($1, $2, $3, $4, $5)
		returning id`

	err := r.pool.QueryRow(ctx, query, order.Status, order.ClientId, order.ItemId, order.ItemCount, order.TotalPrice).Scan(&order.Id)
	if err != nil {
		return entity.Order{}, fmt.Errorf("pool.QueryRow: %w", err)
	}

	return order, nil
}

func (r *PgOrderRepository) Get(ctx context.Context, id entity.OrderID) (entity.Order, error) {
	const query = `
		SELECT id, status, client_id, item_id, item_count, total_price
		FROM orders
		WHERE id = $1`

	var order entity.Order
	err := r.pool.QueryRow(ctx, query, id).Scan(&order.Id, &order.Status, &order.ClientId, &order.ItemId, &order.ItemCount, &order.TotalPrice)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Order{}, model.ErrOrderNotFound
		}
		return entity.Order{}, fmt.Errorf("pool.QueryRow: %w", err)
	}

	return order, nil
}

func (r *PgOrderRepository) ChangeStatus(ctx context.Context, id entity.OrderID, status entity.OrderStatus) error {
	const query = `
		UPDATE orders
		SET status = $1
		WHERE id = $2`

	_, err := r.pool.Exec(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("pool.QueryRow: %w", err)
	}

	return nil
}
