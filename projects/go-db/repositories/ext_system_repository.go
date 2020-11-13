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
INSERT INTO external_systems ("external_system_id", "description", "post_results_url")
VALUES ($1, $2, $3);
`
	insertExtSystemWithoutIDStatement = `
INSERT INTO external_systems ("description", "post_results_url")
VALUES ($1, $2);
`
)

func (repo *ExtSystemRepository) InsertExtSystem(extSystemDAO dao.ExtSystemDAO) error {
	p := repo.DBConn.GetPool()
	ctx := context.Background()
	conn, err := p.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("insert extSystem: acquire connection: %v", err)
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

	row, err := conn.Query(ctx, statement, args...)
	if err != nil {
		return fmt.Errorf("insert extSystem: %v", err)
	}
	row.Close()

	return nil
}

func (repo *ExtSystemRepository) SelectExtSystems() ([]dao.ExtSystemDAO, error) {
	// TODO: for web client need list of extSystems [{id, description}, ...]
	return nil, nil
}
