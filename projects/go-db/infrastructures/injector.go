package infrastructures

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

type IInjector interface {
	CloseStoragesConnections()

	// INFO: controllers
	InjectGameController() controllers.GameController
	InjectGamePrepareController() controllers.GamePrepareController

	InjectExtSystemController() controllers.ExtSystemController

	InjectScreenshotController() controllers.ScreenshotController

	InjectScreenshotSetAnswerController() controllers.ScreenshotSetAnswerController
	InjectScreenshotResultsController() controllers.ScreenshotResultsController

	InjectStatisticUserController() controllers.StatisticUserController
	InjectStatisticLeaderboardController() controllers.StatisticLeaderboardController
	InjectStatisticGameController() controllers.StatisticGameController

	// INFO: services
	InjectGameCacheService() services.GameCacheService

	// TODO: TEST
	InjectImageService() services.ImageService
	InjectValidateFacesService() services.ValidateFacesService
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
	attachSourceToGameService := k.InjectAttachSourceToGameService()
	extSystemService := k.InjectExtSystemService()
	gameService := k.InjectGameService()
	sourceService := k.InjectSourceService()

	return controllers.GameController{
		AttachSourceToGameService: &attachSourceToGameService,
		ExtSystemService:          &extSystemService,
		GameService:               &gameService,
		SourceService:             &sourceService,
	}
}

func (k *kernel) InjectGamePrepareController() controllers.GamePrepareController {
	gameCacheService := k.InjectGameCacheService()

	return controllers.GamePrepareController{GameCacheService: &gameCacheService}
}

func (k *kernel) InjectExtSystemController() controllers.ExtSystemController {
	service := k.InjectExtSystemService()

	return controllers.ExtSystemController{ExtSystemService: &service}
}

func (k *kernel) InjectScreenshotController() controllers.ScreenshotController {
	gameCacheService := k.InjectGameCacheService()
	screenshotCacheService := k.InjectScreenshotCacheService()
	imageService := k.InjectImageService()
	userService := k.InjectUserService()

	return controllers.ScreenshotController{
		GameCacheService:       &gameCacheService,
		ScreenshotCacheService: &screenshotCacheService,
		ImageService:           &imageService,
		UserService:            &userService,
	}
}

func (k *kernel) InjectScreenshotSetAnswerController() controllers.ScreenshotSetAnswerController {
	activeUsersService := k.InjectActiveUsersService()
	screenshotCacheService := k.InjectScreenshotCacheService()
	screenshotUserAnswerService := k.InjectScreenshotUserAnswerService()

	return controllers.ScreenshotSetAnswerController{
		ActiveUsersService:          &activeUsersService,
		ScreenshotCacheService:      &screenshotCacheService,
		ScreenshotUserAnswerService: &screenshotUserAnswerService,
	}
}

func (k *kernel) InjectScreenshotResultsController() controllers.ScreenshotResultsController {
	answerService := k.InjectAnswerService()
	gameService := k.InjectGameService()
	screenshotService := k.InjectScreenshotService()

	return controllers.ScreenshotResultsController{
		AnswerService:     &answerService,
		GameService:       &gameService,
		ScreenshotService: &screenshotService,
	}
}

func (k *kernel) InjectStatisticUserController() controllers.StatisticUserController {
	extSystemService := k.InjectExtSystemService()
	answerService := k.InjectAnswerService()
	gameService := k.InjectGameService()
	userService := k.InjectUserService()
	durationService := k.InjectDurationService()

	return controllers.StatisticUserController{
		GameService:      &gameService,
		ExtSystemService: &extSystemService,
		AnswerService:    &answerService,
		UserService:      &userService,
		DurationService:  &durationService,
	}
}

func (k *kernel) InjectStatisticLeaderboardController() controllers.StatisticLeaderboardController {
	extSystemService := k.InjectExtSystemService()
	gameService := k.InjectGameService()
	durationService := k.InjectDurationService()
	leaderboardService := k.InjectLeaderboardService()

	return controllers.StatisticLeaderboardController{
		ExtSystemService:   &extSystemService,
		GameService:        &gameService,
		DurationService:    &durationService,
		LeaderboardService: &leaderboardService,
	}
}

func (k *kernel) InjectStatisticGameController() controllers.StatisticGameController {
	extSystemService := k.InjectExtSystemService()
	gameService := k.InjectGameService()
	statisticGameService := k.InjectStatisticGameService()

	return controllers.StatisticGameController{
		ExtSystemService:     &extSystemService,
		GameService:          &gameService,
		StatisticGameService: &statisticGameService,
	}
}
