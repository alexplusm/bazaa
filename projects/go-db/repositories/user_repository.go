package repositories

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type UserRepository struct {
	DBConn interfaces.IDBHandler
}

// ON CONFLICT DO NOTHING
const (
	insertUserStatement = `INSERT INTO users ("user_id") VALUES ($1);`
)

func (repo *UserRepository) InsertUser(user dao.UserDAO) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("insert user: acquire connection: %v", err)
	}
	defer conn.Release()

	row, err := conn.Query(context.Background(), insertUserStatement, user.UserID)
	if err != nil {
		return fmt.Errorf("insert user: %v", err)
	}
	row.Close()

	return nil
}
