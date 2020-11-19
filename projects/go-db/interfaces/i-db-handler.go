package interfaces

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type IDBHandler interface {
	GetPool() *pgxpool.Pool
}
