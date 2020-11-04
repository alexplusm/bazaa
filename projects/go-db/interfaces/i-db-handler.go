package interfaces

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO: create another package? "core/interfaces"
// "core/utils"

type IDBHandler interface {
	GetConn() *pgxpool.Pool // TODO: late encapsulate this in methods (below)
	// Execute(statement string)
	// Query(statement string) (IRow, error)
	// ...
}
