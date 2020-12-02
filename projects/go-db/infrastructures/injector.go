package infrastructures

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/repositories"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

type IInjector interface {
	CloseStoragesConnections()

	// INFO: controllers
	InjectGameCreateController() controllers.GameCreateController
	InjectGameUpdateController() controllers.GameUpdateController
	InjectGameListController() controllers.GameListController
	InjectGamePrepareController() controllers.GamePrepareController
	InjectGameInfoController() controllers.GameInfoController

	InjectExtSystemController() controllers.ExtSystemController

	InjectScreenshotGetController() controllers.ScreenshotGetController
	InjectScreenshotSetAnswerController() controllers.ScreenshotSetAnswerController
	InjectScreenshotResultsController() controllers.ScreenshotResultsController

	InjectStatisticsUserController() controllers.StatisticsUserController
	InjectStatisticsLeaderboardController() controllers.StatisticsLeaderboardController
	InjectStatisticsGameController() controllers.StatisticsGameController

	// INFO: services
	InjectGameCacheService() services.GameCacheService
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
				// TODO: try to reconnect?
				err = fmt.Errorf("injector: database connection: %v", pqslErr)
			}
			redisClient := initRedis()
			k = &kernel{pool, redisClient}
		})
	}
	return k, err
}

func (k *kernel) CloseStoragesConnections() {
	k.pool.Close()
	if err := k.redisClient.Close(); err != nil {
		log.Error("redis: error while close connection: ", err)
	}
}

// INFO: Controllers

func (k *kernel) InjectGameCreateController() controllers.GameCreateController {
	handler := &PSQLHandler{k.pool}

	repo := &repositories.GameRepository{DBConn: handler}
	service := &services.GameService{GameRepo: repo}
	controller := controllers.GameCreateController{GameService: service}

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

func (k *kernel) InjectGameListController() controllers.GameListController {
	handler := &PSQLHandler{k.pool}

	gameRepo := &repositories.GameRepository{DBConn: handler}
	extSystemRepo := &repositories.ExtSystemRepository{DBConn: handler}

	gameService := &services.GameService{GameRepo: gameRepo}
	extSystemService := &services.ExtSystemService{ExtSystemRepo: extSystemRepo}

	controller := controllers.GameListController{
		GameService: gameService, ExtSystemService: extSystemService,
	}

	return controller
}

func (k *kernel) InjectGamePrepareController() controllers.GamePrepareController {
	gameCacheService := k.InjectGameCacheService()
	controller := controllers.GamePrepareController{GameCacheService: &gameCacheService}

	return controller
}

func (k *kernel) InjectGameInfoController() controllers.GameInfoController {
	handler := &PSQLHandler{k.pool}

	sourceRepo := &repositories.SourceRepository{DBConn: handler}
	gameRepo := &repositories.GameRepository{DBConn: handler}
	gameService := &services.GameService{GameRepo: gameRepo}
	sourceService := &services.SourceService{SourceRepo: sourceRepo}

	controller := controllers.GameInfoController{
		GameService: gameService, SourceService: sourceService,
	}

	return controller
}

func (k *kernel) InjectExtSystemController() controllers.ExtSystemController {
	service := k.InjectExtSystemService()
	return controllers.ExtSystemController{ExtSystemService: &service}
}

func (k *kernel) InjectScreenshotGetController() controllers.ScreenshotGetController {
	redisHandler := &RedisHandler{k.redisClient}
	dbHandler := &PSQLHandler{k.pool}

	screenshotRepo := &repositories.ScreenshotRepository{DBConn: dbHandler}
	gameRepo := &repositories.GameRepository{DBConn: dbHandler}
	userRepo := &repositories.UserRepository{DBConn: dbHandler}

	screenshotCacheService := &services.ScreenshotCacheService{RedisClient: redisHandler}
	gameCacheService := &services.GameCacheService{
		RedisClient: redisHandler, ScreenshotRepo: screenshotRepo, GameRepo: gameRepo,
	}
	userService := &services.UserService{UserRepo: userRepo}

	controller := controllers.ScreenshotGetController{
		ScreenshotCacheService: screenshotCacheService,
		GameCacheService:       gameCacheService,
		UserService:            userService,
	}

	return controller
}

func (k *kernel) InjectScreenshotSetAnswerController() controllers.ScreenshotSetAnswerController {
	redisHandler := &RedisHandler{k.redisClient}
	dbHandler := &PSQLHandler{k.pool}

	answerRepo := &repositories.AnswerRepository{DBConn: dbHandler}
	screenshotRepo := &repositories.ScreenshotRepository{DBConn: dbHandler}

	screenshotCacheService := &services.ScreenshotCacheService{RedisClient: redisHandler}
	screenshotUserAnswerService := &services.ScreenshotUserAnswerService{
		AnswerRepo: answerRepo, ScreenshotRepo: screenshotRepo,
	}
	activeUsersService := &services.ActiveUsersService{RedisClient: redisHandler}

	controller := controllers.ScreenshotSetAnswerController{
		ScreenshotCacheService:      screenshotCacheService,
		ScreenshotUserAnswerService: screenshotUserAnswerService,
		ActiveUsersService:          activeUsersService,
	}

	return controller
}

func (k *kernel) InjectScreenshotResultsController() controllers.ScreenshotResultsController {
	handler := &PSQLHandler{k.pool}

	answerRepo := &repositories.AnswerRepository{DBConn: handler}
	gameRepo := &repositories.GameRepository{DBConn: handler}
	screenshotRepo := &repositories.ScreenshotRepository{DBConn: handler}

	answerService := &services.AnswerService{AnswerRepo: answerRepo}
	gameService := &services.GameService{GameRepo: gameRepo}
	screenshotService := &services.ScreenshotService{ScreenshotRepo: screenshotRepo}

	controller := controllers.ScreenshotResultsController{
		AnswerService: answerService, GameService: gameService,
		ScreenshotService: screenshotService,
	}

	return controller
}

func (k *kernel) InjectStatisticsUserController() controllers.StatisticsUserController {
	handler := &PSQLHandler{k.pool}

	answerRepo := &repositories.AnswerRepository{DBConn: handler}
	extSystemRepo := &repositories.ExtSystemRepository{DBConn: handler}
	gameRepo := &repositories.GameRepository{DBConn: handler}
	userRepo := &repositories.UserRepository{DBConn: handler}

	answerService := &services.AnswerService{AnswerRepo: answerRepo}
	extSystemService := &services.ExtSystemService{ExtSystemRepo: extSystemRepo}
	gameService := &services.GameService{GameRepo: gameRepo}
	userService := &services.UserService{UserRepo: userRepo}

	controller := controllers.StatisticsUserController{
		GameService: gameService, ExtSystemService: extSystemService,
		AnswerService: answerService, UserService: userService,
	}

	return controller
}

func (k *kernel) InjectStatisticsLeaderboardController() controllers.StatisticsLeaderboardController {
	handler := &PSQLHandler{k.pool}

	extSystemRepo := &repositories.ExtSystemRepository{DBConn: handler}
	extSystemService := &services.ExtSystemService{ExtSystemRepo: extSystemRepo}

	answerRepo := &repositories.AnswerRepository{DBConn: handler}
	answerService := &services.AnswerService{AnswerRepo: answerRepo}

	gameRepo := &repositories.GameRepository{DBConn: handler}
	gameService := &services.GameService{GameRepo: gameRepo}

	controller := controllers.StatisticsLeaderboardController{
		ExtSystemService: extSystemService, GameService: gameService,
		AnswerService: answerService,
	}

	return controller
}

func (k *kernel) InjectStatisticsGameController() controllers.StatisticsGameController {
	handler := &PSQLHandler{k.pool}
	redisHandler := &RedisHandler{k.redisClient}

	extSystemRepo := &repositories.ExtSystemRepository{DBConn: handler}
	extSystemService := &services.ExtSystemService{ExtSystemRepo: extSystemRepo}

	gameRepo := &repositories.GameRepository{DBConn: handler}
	gameService := &services.GameService{GameRepo: gameRepo}

	screenshotRepo := &repositories.ScreenshotRepository{DBConn: handler}
	screenshotService := &services.ScreenshotService{ScreenshotRepo: screenshotRepo}

	answerRepo := &repositories.AnswerRepository{DBConn: handler}
	answerService := &services.AnswerService{AnswerRepo: answerRepo}

	activeUsersService := &services.ActiveUsersService{RedisClient: redisHandler}

	controller := controllers.StatisticsGameController{
		ExtSystemService: extSystemService, GameService: gameService,
		ScreenshotService: screenshotService, AnswerService: answerService,
		ActiveUsersService: activeUsersService,
	}

	return controller
}

// INFO: services

func (k *kernel) InjectGameCacheService() services.GameCacheService {
	redisHandler := &RedisHandler{k.redisClient}
	dbHandler := &PSQLHandler{k.pool}

	screenshotRepo := &repositories.ScreenshotRepository{DBConn: dbHandler}
	gameRepo := &repositories.GameRepository{DBConn: dbHandler}
	service := services.GameCacheService{
		RedisClient: redisHandler, ScreenshotRepo: screenshotRepo, GameRepo: gameRepo,
	}

	return service
}

func (k *kernel) InjectExtSystemService() services.ExtSystemService {
	handler := &PSQLHandler{k.pool}
	repo := &repositories.ExtSystemRepository{DBConn: handler}
	service := services.ExtSystemService{ExtSystemRepo: repo}

	return service
}
