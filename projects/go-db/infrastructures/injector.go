package infrastructures

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/repositories"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

type IInjector interface {
	CloseStoragesConnections()

	// INFO: controllers
	InjectCreateGameController() controllers.CreateGameController
	InjectUpdateGameController() controllers.UpdateGameController
}

func (k *kernel) InjectCreateGameController() controllers.CreateGameController {
	handler := &PSQLHandler{k.pool} // TODO: in kernel? | after creation end point for game creation

	repo := &repositories.GameRepository{handler}
	service := &services.CreateGameService{repo}
	controller := controllers.CreateGameController{service}

	return controller
}

func (k *kernel) InjectUpdateGameController() controllers.UpdateGameController {
	handler := &PSQLHandler{k.pool}

	gameRepo := &repositories.GameRepository{handler}
	sourceRepo := &repositories.SourceRepository{handler}
	screenshotRepo := &repositories.ScreenshotRepository{handler}
	service := &services.UpdateGameService{
		GameRepo: gameRepo, SourceRepo: sourceRepo, ScreenshotRepo: screenshotRepo,
	}
	controller := controllers.UpdateGameController{service}

	return controller
}

func (k *kernel) CloseStoragesConnections() {
	// TODO: test
	// TODO: Redis Close
	// TODO: create dumps?
	k.pool.Close()
	k.redisClient.Close() // TODO:error
}

type kernel struct {
	pool        *pgxpool.Pool
	redisClient *redis.Client
}

var (
	k         *kernel
	singleton sync.Once
)

func Injector() IInjector {
	if k == nil {
		singleton.Do(func() {
			pool, err := initPostgresql()
			if err != nil {
				// todo: need try to reconnect? how undo "singleton"?
				fmt.Println("Error connection")
			}

			redisClient := initRedis()

			k = &kernel{pool, redisClient}
		})
	}

	return k
}
