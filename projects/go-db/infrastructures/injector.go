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
	InjectGameController() controllers.GameController
	InjectGamePrepareController() controllers.GamePrepareController

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

func (k *kernel) InjectGameController() controllers.GameController {
	gameService := k.InjectGameService()
	extSystemService := k.InjectExtSystemService()
	sourceService := k.InjectSourceService()
	attachSourceToGameService := k.InjectAttachSourceToGameService()
	return controllers.GameController{
		GameService: &gameService, ExtSystemService: &extSystemService,
		SourceService: &sourceService, AttachSourceToGameService: &attachSourceToGameService,
	}
}

func (k *kernel) InjectGamePrepareController() controllers.GamePrepareController {
	gameCacheService := k.InjectGameCacheService()
	controller := controllers.GamePrepareController{GameCacheService: &gameCacheService}

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

	screenshotCacheService := &services.ScreenshotCacheService{RedisClient: redisHandler}
	gameCacheService := &services.GameCacheService{
		RedisClient: redisHandler, ScreenshotRepo: screenshotRepo, GameRepo: gameRepo,
	}

	userService := k.InjectUserService()
	imageService := k.InjectImageService()

	controller := controllers.ScreenshotGetController{
		ScreenshotCacheService: screenshotCacheService,
		GameCacheService:       gameCacheService,
		UserService:            &userService,
		ImageService:           &imageService,
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

	durationService := k.InjectDurationService()

	controller := controllers.StatisticsUserController{
		GameService: gameService, ExtSystemService: extSystemService,
		AnswerService: answerService, UserService: userService,
		DurationService: &durationService,
	}

	return controller
}

func (k *kernel) InjectStatisticsLeaderboardController() controllers.StatisticsLeaderboardController {
	extSystemService := k.InjectExtSystemService()
	gameService := k.InjectGameService()
	durationService := k.InjectDurationService()
	leaderboardService := k.InjectLeaderboardService()

	return controllers.StatisticsLeaderboardController{
		ExtSystemService: &extSystemService, GameService: &gameService,
		DurationService: &durationService, LeaderboardService: &leaderboardService,
	}
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
	return services.ExtSystemService{ExtSystemRepo: repo}
}

func (k *kernel) InjectDurationService() services.DurationService {
	return services.DurationService{}
}

func (k *kernel) InjectUserService() services.UserService {
	handler := &PSQLHandler{k.pool}
	repo := &repositories.UserRepository{DBConn: handler}
	return services.UserService{UserRepo: repo}
}

func (k *kernel) InjectGameService() services.GameService {
	handler := &PSQLHandler{k.pool}
	repo := &repositories.GameRepository{DBConn: handler}
	return services.GameService{GameRepo: repo}
}

func (k *kernel) InjectSourceService() services.SourceService {
	handler := &PSQLHandler{k.pool}
	repo := &repositories.SourceRepository{DBConn: handler}
	return services.SourceService{SourceRepo: repo}
}

func (k *kernel) InjectAnswerService() services.AnswerService {
	handler := &PSQLHandler{k.pool}
	repo := &repositories.AnswerRepository{DBConn: handler}
	return services.AnswerService{AnswerRepo: repo}
}

func (k *kernel) InjectAttachSourceToGameService() services.AttachSourceToGameService {
	handler := &PSQLHandler{k.pool}
	gameRepo := &repositories.GameRepository{DBConn: handler}
	sourceRepo := &repositories.SourceRepository{DBConn: handler}
	screenshotRepo := &repositories.ScreenshotRepository{DBConn: handler}
	fileService := k.InjectFileService()

	return services.AttachSourceToGameService{
		GameRepo: gameRepo, SourceRepo: sourceRepo, ScreenshotRepo: screenshotRepo,
		FileService: &fileService,
	}
}

func (k *kernel) InjectLeaderboardService() services.LeaderboardService {
	answerService := k.InjectAnswerService()
	return services.LeaderboardService{AnswerService: &answerService}
}

func (k *kernel) InjectImageService() services.ImageService {
	return services.ImageService{}
}

func (k *kernel) InjectFileService() services.FileService {
	return services.FileService{}
}
