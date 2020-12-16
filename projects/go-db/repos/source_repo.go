package repos

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type SourceRepo struct {
	DBConn interfaces.IDBHandler
}

const (
	insertSourceStatement = `
INSERT INTO sources ("game_id", "type", "created_at", "value")
VALUES ($1, $2, $3, $4)
RETURNING "source_id"
`
	selectSourcesByGameStatement = `
SELECT "source_id", "game_id", "type", "created_at", "value"
FROM sources
WHERE sources.game_id = ($1)
`
)

func (repo *SourceRepo) InsertOne(source dao.SourceInsertDAO) (string, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("insert source: acquire connection: %v", err)
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		insertSourceStatement,
		source.GameID, source.Type, source.CreatedAt, source.Value,
	)

	var sourceID string

	err = row.Scan(&sourceID)
	if err != nil {
		return "", fmt.Errorf("insert source: query: %v", err)
	}

	return sourceID, nil
}

func (repo *SourceRepo) SelectListByGame(gameID string) ([]dao.SourceRetrieveDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select sources by game: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(
		context.Background(), selectSourcesByGameStatement, gameID,
	)
	if err != nil {
		return nil, fmt.Errorf("select sources by game: %v", err)
	}
	defer rows.Close()

	list := make([]dao.SourceRetrieveDAO, 0, 10)

	for rows.Next() {
		s := dao.SourceRetrieveDAO{}
		err = rows.Scan(&s.SourceID, &s.GameID, &s.Type, &s.CreatedAt, &s.Value)
		if err != nil {
			return nil, fmt.Errorf("select sources by game: %v", err)
		}
		list = append(list, s)
	}

	return list, nil
}
