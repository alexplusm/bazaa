package repositories

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

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
	selectExtSystems = `
SELECT "ext_system_id", "description", "post_results_url"
FROM ext_systems;
`
)

func (repo *ExtSystemRepository) InsertOne(
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

func (repo *ExtSystemRepository) SelectList() ([]dao.ExtSystemDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("extSystem list: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), selectExtSystems)

	if err != nil {
		return nil, fmt.Errorf("extSystem list: %v", err)
	}
	defer rows.Close()

	list := make([]dao.ExtSystemDAO, 0, 1024)

	for rows.Next() {
		es := dao.ExtSystemDAO{}
		err = rows.Scan(&es.ID, &es.Description, &es.PostResultsURL)
		if err != nil {
			log.Error("extSystem list: ", err)
			continue
		}
		list = append(list, es)
	}
	if rows.Err() != nil {
		log.Error("extSystem list: ", rows.Err())
	}

	return list, nil
}

func (repo *ExtSystemRepository) Exist(extSystemID string) (bool, error) {
	p := repo.DBConn.GetPool()
	ctx := context.Background()
	conn, err := p.Acquire(ctx)
	if err != nil {
		return false, fmt.Errorf("extSystem exist: acquire connection: %v", err)
	}
	defer conn.Release()

	var count int64

	row := conn.QueryRow(ctx, existExtSystemStatement, extSystemID)
	if row.Scan(&count) != nil {
		return false, fmt.Errorf("extSystem exist: %v", err)
	}

	return count != 0, nil
}
