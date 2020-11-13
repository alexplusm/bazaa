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

	// INFO: services
	InjectGameCacheService() services.GameCacheService
}

func (k *kernel) InjectGameCreateController() controllers.GameCreateController {
	handler := &PSQLHandler{k.pool}

	repo := &repositories.GameRepository{DBConn: handler}
	service := &services.GameService{GameRepo: repo}
	controller := controllers.GameCreateController{Service: service}

	return controller
}

func (k *kernel) InjectGameUpdateController() controllers.GameUpdateController {
	handler := &PSQLHandler{k.pool}

	gameRepo := &repositories.GameRepository{DBConn: handler}
	sourceRepo := &repositories.SourceRepository{DBConn: handler}
	screenshotRepo := &repositories.ScreenshotRepository{DBConn: handler}

	gameService := &services.GameService{GameRepo: gameRepo}
	attachSourceToGameService := &services.AttachSourceToGameService{
		GameRepo: gameRepo, SourceRepo: sourceRepo, ScreenshotRepo: screenshotRepo,
	}

	controller := controllers.GameUpdateController{
		GameService: gameService, AttachSourceToGameService: attachSourceToGameService,
	}

	return controller
}

func (k *kernel) InjectExtSystemCreateController() controllers.ExtSystemCreateController {
	handler := &PSQLHandler{k.pool}

	repo := &repositories.ExtSystemRepository{DBConn: handler}
	service := &services.ExtSystemService{ExtSystemRepo: repo}
	controller := controllers.ExtSystemCreateController{ExtSystemService: service}

	return controller
}

func (k *kernel) InjectScreenshotGetController() controllers.ScreenshotGetController {
	controller := controllers.ScreenshotGetController{}

	return controller
}

func (k *kernel) InjectGameCacheService() services.GameCacheService {
	redisHandler := &RedisHandler{k.redisClient}
	DBhandler := &PSQLHandler{k.pool}

	screenshotRepo := &repositories.ScreenshotRepository{DBConn: DBhandler}
	gameRepo := &repositories.GameRepository{DBConn: DBhandler}
	service := services.GameCacheService{
		RedisClient: redisHandler, ScreenshotRepo: screenshotRepo, GameRepo: gameRepo,
	}

	return service
}

func (k *kernel) CloseStoragesConnections() {
	// TODO: create dumps?
	k.pool.Close()
	if err := k.redisClient.Close(); err != nil {
		fmt.Printf("redis: error while close connection: %v\n", err)
	}
}

type kernel struct {
	pool        *pgxpool.Pool
	redisClient *redis.Client
}

var (
	k         *kernel
	singleton sync.Once
)

func Injector() (IInjector, error) {
	var err error = nil
	if k == nil {
		singleton.Do(func() {
			pool, pqslErr := initPostgresql()
			if pqslErr != nil {
				// TODO: try to reconnect? need patterns
				err = fmt.Errorf("injector: database connection: %v", pqslErr)
			}
			redisClient := initRedis()
			k = &kernel{pool, redisClient}
		})
	}
	return k, err
}
