package repositories

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ExtSystemRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertExtSystemWithIDStatement = `
INSERT INTO ext_systems ("ext_system_id", "description", "post_results_url")
VALUES ($1, $2, $3)
RETURNING "ext_system_id";
`
	insertExtSystemWithoutIDStatement = `
INSERT INTO ext_systems ("description", "post_results_url")
VALUES ($1, $2)
RETURNING "ext_system_id";
`
	existExtSystemStatement = `
SELECT COUNT(1)
FROM ext_systems
WHERE "ext_system_id" = ($1);
`
)

func (repo *ExtSystemRepository) InsertExtSystem(
	extSystemDAO dao.ExtSystemDAO,
) (string, error) {
	p := repo.DBConn.GetPool()
	ctx := context.Background()
	conn, err := p.Acquire(ctx)
	if err != nil {
		return "", fmt.Errorf("insert extSystem: acquire connection: %v", err)
	}
	defer conn.Release()

	var args []interface{}
	var statement string

	if extSystemDAO.HasID() {
		statement = insertExtSystemWithIDStatement
		args = []interface{}{extSystemDAO.ID, extSystemDAO.Description, extSystemDAO.PostResultsURL}
	} else {
		statement = insertExtSystemWithoutIDStatement
		args = []interface{}{extSystemDAO.Description, extSystemDAO.PostResultsURL}
	}

	row := conn.QueryRow(ctx, statement, args...)

	var extSystemID string
	if err = row.Scan(&extSystemID); err != nil {
		return "", fmt.Errorf("insert extSystem: %v", err)
	}

	return extSystemID, nil
}

func (repo *ExtSystemRepository) SelectExtSystems() ([]dao.ExtSystemDAO, error) {
	// TODO: for web client need list of extSystems [{id, description}, ...]
	return nil, nil
}

func (repo *ExtSystemRepository) ExtSystemExist(extSystemID string) (bool, error) {
	p := repo.DBConn.GetPool()
	ctx := context.Background()
	conn, err := p.Acquire(ctx)
	if err != nil {
		return false, fmt.Errorf("extSystem exist: acquire connection: %v", err)
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, existExtSystemStatement, extSystemID)

	var count int64

	if row.Scan(&count) != nil {
		return false, fmt.Errorf("extSystem exist: %v", err)
	}

	fmt.Println("COUNT: ", count, extSystemID)

	return count != 0, nil
}
