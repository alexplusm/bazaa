package repositories

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type SourceRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertSourceStatement = `
INSERT INTO sources ("game_id", "source_type", "created_at")
VALUES ($1, $2, $3)
RETURNING "source_id";
`
	// TODO: for schedule
	insertSourceWithExistingIDStatement = `
INSERT INTO sources ("game_id", "source_id", "source_type", "created_at")
VALUES ($1, $2, $3, $4)
RETURNING "source_id";
`
)

func (repo *SourceRepository) InsertSource(source dao.SourceDAO) (string, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("insert source: acquire connection: %v", err)
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		insertSourceStatement,
		source.GameID, source.Type, source.CreatedAt,
	)

	var sourceID string

	err = row.Scan(&sourceID)
	if err != nil {
		return "", fmt.Errorf("insert source: query: %v", err)
	}

	return sourceID, nil
}
