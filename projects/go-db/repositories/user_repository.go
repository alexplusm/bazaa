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

const (
	insertUserStatement = `
INSERT INTO users ("user_id") 
VALUES ($1)
`
	existUserStatement = `
SELECT COUNT(1)
FROM users
WHERE "user_id" = ($1)
`
)

func (repo *UserRepository) InsertOne(user dao.UserDAO) error {
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

func (repo *UserRepository) Exist(userID string) (bool, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return false, fmt.Errorf("user exist: acquire connection: %v", err)
	}
	defer conn.Release()

	var count int64

	row := conn.QueryRow(context.Background(), existUserStatement, userID)
	if row.Scan(&count) != nil {
		return false, fmt.Errorf("user exist: %v", err)
	}

	return count != 0, nil
}
