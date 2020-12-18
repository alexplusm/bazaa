package infrastructures

import (
	"github.com/Alexplusm/bazaa/projects/go-db/repos"
	"github.com/Alexplusm/bazaa/projects/go-db/services"
)

func (k *kernel) InjectGameCacheService() services.GameCacheService {
	redisHandler := &RedisHandler{k.redisClient}
	dbHandler := &PSQLHandler{k.pool}
	screenshotRepo := &repos.ScreenshotRepo{DBConn: dbHandler}
	gameRepo := &repos.GameRepo{DBConn: dbHandler}

	return services.GameCacheService{
		RedisClient: redisHandler, ScreenshotRepo: screenshotRepo, GameRepo: gameRepo,
	}
}

func (k *kernel) InjectExtSystemService() services.ExtSystemService {
	handler := &PSQLHandler{k.pool}
	repo := &repos.ExtSystemRepo{DBConn: handler}

	return services.ExtSystemService{ExtSystemRepo: repo}
}

func (k *kernel) InjectDurationService() services.DurationService {
	return services.DurationService{}
}

func (k *kernel) InjectUserService() services.UserService {
	handler := &PSQLHandler{k.pool}
	repo := &repos.UserRepo{DBConn: handler}

	return services.UserService{UserRepo: repo}
}

func (k *kernel) InjectGameService() services.GameService {
	handler := &PSQLHandler{k.pool}
	repo := &repos.GameRepo{DBConn: handler}

	return services.GameService{GameRepo: repo}
}

func (k *kernel) InjectSourceService() services.SourceService {
	handler := &PSQLHandler{k.pool}
	repo := &repos.SourceRepo{DBConn: handler}

	return services.SourceService{SourceRepo: repo}
}

func (k *kernel) InjectAnswerService() services.AnswerService {
	handler := &PSQLHandler{k.pool}
	repo := &repos.AnswerRepo{DBConn: handler}

	return services.AnswerService{AnswerRepo: repo}
}

func (k *kernel) InjectAttachSourceToGameService() services.AttachSourceToGameService {
	handler := &PSQLHandler{k.pool}
	gameRepo := &repos.GameRepo{DBConn: handler}
	screenshotRepo := &repos.ScreenshotRepo{DBConn: handler}
	fileService := k.InjectFileService()
	sourceService := k.InjectSourceService()
	imageFilterService := k.InjectImageFilterService()

	return services.AttachSourceToGameService{
		GameRepo:           gameRepo,
		ScreenshotRepo:     screenshotRepo,
		SourceService:      &sourceService,
		FileService:        &fileService,
		ImageFilterService: &imageFilterService,
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

func (k *kernel) InjectScreenshotService() services.ScreenshotService {
	handler := &PSQLHandler{k.pool}
	screenshotRepo := &repos.ScreenshotRepo{DBConn: handler}

	return services.ScreenshotService{ScreenshotRepo: screenshotRepo}
}

func (k *kernel) InjectActiveUsersService() services.ActiveUsersService {
	redisHandler := &RedisHandler{k.redisClient}
	cacheKeyService := k.InjectCacheKeyService()

	return services.ActiveUsersService{
		RedisClient:     redisHandler,
		CacheKeyService: &cacheKeyService,
	}
}

func (k *kernel) InjectStatisticGameService() services.StatisticGameService {
	screenshotService := k.InjectScreenshotService()
	answerService := k.InjectAnswerService()
	activeUsersService := k.InjectActiveUsersService()

	return services.StatisticGameService{
		ActiveUsersService: &activeUsersService,
		AnswerService:      &answerService,
		ScreenshotService:  &screenshotService,
	}
}

func (k *kernel) InjectScreenshotCacheService() services.ScreenshotCacheService {
	redisHandler := &RedisHandler{k.redisClient}
	cacheKeyService := k.InjectCacheKeyService()

	return services.ScreenshotCacheService{RedisClient: redisHandler, CacheKeyService: &cacheKeyService}
}

func (k *kernel) InjectScreenshotUserAnswerService() services.ScreenshotUserAnswerService {
	dbHandler := &PSQLHandler{k.pool}
	answerRepo := &repos.AnswerRepo{DBConn: dbHandler}
	screenshotRepo := &repos.ScreenshotRepo{DBConn: dbHandler}

	return services.ScreenshotUserAnswerService{
		AnswerRepo:     answerRepo,
		ScreenshotRepo: screenshotRepo,
	}
}

func (k *kernel) InjectValidateFacesService() services.ValidateFacesService {
	return services.ValidateFacesService{}
}

func (k *kernel) InjectImageFilterService() services.ImageFilterService {
	validateFacesService := k.InjectValidateFacesService()
	imageService := k.InjectImageService()

	return services.ImageFilterService{
		ValidateFacesService: &validateFacesService,
		ImageService:         &imageService,
	}
}

func (k *kernel) InjectCacheKeyService() services.CacheKeyService {
	return services.CacheKeyService{}
}
