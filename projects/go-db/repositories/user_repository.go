package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UserRepository struct {
	DBConn interfaces.IDBHandler
}
