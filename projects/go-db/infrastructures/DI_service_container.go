package infrastructures

import (
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/repositories"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

type IServiceContainer interface {
	// INFO: controllers
	InjectCreateGameController() controllers.CreateGameController

	CloseStoragesConnections()
}

func (k *kernel) InjectCreateGameController() controllers.CreateGameController {
	handler := &PSQLHandler{k.pool} // TODO: in kernel?

	repo := &repositories.GameRepository{handler}
	service := &services.GameService{repo}
	controller := controllers.CreateGameController{service}

	return controller
}

func (k *kernel) CloseStoragesConnections() {
	// TODO: test
	// TODO: Redis Close
	// TODO: create dumps?
	k.pool.Close()
}

type kernel struct {
	pool *pgxpool.Pool
}

var (
	k         *kernel
	singleton sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		singleton.Do(func() {
			pool, err := initPostgresql()
			if err != nil {
				// todo: need try to reconnect? how undo "singleton"?
				fmt.Println("Error connection")
			}
			k = &kernel{pool}
		})
	}

	return k
}
