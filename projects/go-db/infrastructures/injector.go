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
	InjectGameCreateController() controllers.GameCreateController
	InjectGameUpdateController() controllers.GameUpdateController
	InjectExtSystemCreateController() controllers.ExtSystemCreateController
	InjectScreenshotGetController() controllers.ScreenshotGetController

	// TODO: Test: TODO: ServiceInjector and ControllerInjector????
	InjectGameCacheService() services.GameCacheService
}

func (k *kernel) InjectGameCreateController() controllers.GameCreateController {
	handler := &PSQLHandler{k.pool} // TODO: in kernel? | after creation end point for game creation

	repo := &repositories.GameRepository{handler}
	service := &services.GameService{repo}
	controller := controllers.GameCreateController{service}

	return controller
}

func (k *kernel) InjectGameUpdateController() controllers.GameUpdateController {
	handler := &PSQLHandler{k.pool}

	gameRepo := &repositories.GameRepository{handler}
	sourceRepo := &repositories.SourceRepository{handler}
	screenshotRepo := &repositories.ScreenshotRepository{handler}

	gameService := &services.GameService{gameRepo}
	attachSourceToGameService := &services.AttachSourceToGameService{
		GameRepo: gameRepo, SourceRepo: sourceRepo, ScreenshotRepo: screenshotRepo,
	}

	controller := controllers.GameUpdateController{
		gameService, attachSourceToGameService,
	}

	return controller
}

func (k *kernel) InjectExtSystemCreateController() controllers.ExtSystemCreateController {
	handler := &PSQLHandler{k.pool}

	repo := &repositories.ExtSystemRepository{handler}
	service := &services.ExtSystemService{repo}
	controller := controllers.ExtSystemCreateController{service}

	return controller
}

func (k *kernel) InjectScreenshotGetController() controllers.ScreenshotGetController {
	controller := controllers.ScreenshotGetController{}

	return controller
}

func (k *kernel) InjectGameCacheService() services.GameCacheService {
	redisHandler := &RedisHandler{k.redisClient}
	DBhandler := &PSQLHandler{k.pool}

	screenshotRepo := &repositories.ScreenshotRepository{DBhandler}
	service := services.GameCacheService{redisHandler, screenshotRepo}

	return service
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
