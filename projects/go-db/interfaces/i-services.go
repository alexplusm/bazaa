package interfaces

import (
	"archive/zip"
	"mime/multipart"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type IGameService interface {
	Create(game bo.GameBO) (string, error)
	Retrieve(gameID string) (bo.GameBO, error)
	List(extSystemID string) ([]bo.GameBO, error)
	Exist(gameID string) (bool, error)
	FilterGames(gamesID []string, games []bo.GameBO) []bo.GameBO
	GetEarliestGame(games []bo.GameBO) bo.GameBO
}

type IAttachSourceToGameService interface {
	AttachArchives(gameID string, archives []*multipart.FileHeader) error
	AttachSchedules(gameID string) error
	AttachGameResults(gameID string, params bo.AttachGameParams) error
}

type IExtSystemService interface {
	Create(extSystem bo.ExtSystemBO) (string, error)
	Exist(extSystemID string) (bool, error)
	List() ([]bo.ExtSystemBO, error)
}

type IGameCacheService interface {
	PrepareGame(gameID string) error
	GameWithSameExtSystemIDExist(gameID, extSystemID string) bool
}

type IScreenshotCacheService interface {
	CanSetUserAnswerToScreenshot(userID, screenshotID string) bool
	GetScreenshot(gameID, userID string) (dao.ScreenshotURLDAO, bool)
	GetUsersAnswers(screenshotID string) []bo.UserAnswerCacheBO
	SetUserAnswer(userID, screenshotID, answer string) ([]bo.UserAnswerCacheBO, error)
	SetUserAnswerToScreenshot(userID, screenshotID, answer string)
	ScreenshotExist(screenshotID string) bool
	RemoveScreenshot(gameID, screenshotID string)
}

type IScreenshotUserAnswerService interface {
	BuildUserAnswerResponse(
		userID string, answersBO []bo.UserAnswerCacheBO,
	) dto.UserAnswerResponseData
	ScreenshotIsFinished(answers []bo.UserAnswerCacheBO) bool
	SaveUsersAnswers(answers []bo.UserAnswerCacheBO, gameID, screenshotID string)
}

type IUserService interface {
	CreateUser(userID string) error
	UserExist(userID string) (bool, error)
}

type IAnswerService interface {
	GetUserStatistics(
		userID string, gameIDs []string, from, to time.Time,
	) ([]bo.StatisticAnswersDateSlicedBO, error)
	GetScreenshotResults(gameID, screenshotID string) ([]dto.UserAnswerForScreenshotResultDTO, error)
	GetUsersAndScreenshotCountByGame(gameID string) (dao.AnsweredScreenshotsDAO, error)
	ABC(gameID string, from, to time.Time) ([]dao.AnswerScreenshotRetrieveDAO, error) // TODO: BO?
}

type IScreenshotService interface {
	Exist(screenshotID string) (bool, error)
	ScreenshotCountByGame(gameID string) (int, error)
}

type IStatisticGameService interface {
	GetStatistics(games []bo.GameBO) (dto.StatisticGameDTO, error)
}

type ISourceService interface {
	Create(gameId, value string, sourceType int) (string, error)
	ListByGame(gameID string) ([]bo.SourceBO, error)
	GameHasSomeSourceGameId(gameId, sourceGameId string) (bool, error)
}

type IActiveUsersService interface {
	SetUserActivity(gameID, userID string)
	CountOfActiveUsers(gameID string) (int, error)
}

type IDurationService interface {
	GetDurationByGame(from, to string, game bo.GameBO) (time.Time, time.Time, error)
}

type ILeaderboardService interface {
	GetLeaderboard(gameIDs []string, from, to time.Time, limit int) (dto.LeadersResponseDTO, error)
}

type IImageService interface {
	BuildImageURL(imageName string) (string, error)
}

type IFileService interface {
	SaveFiles(files []*multipart.FileHeader, copyPath string) ([]string, error)
	UnzipArchives(archivesPath []string, dstPath string) ([]zip.File, error)
}

type IImageFilterService interface {
	Filter()
}

type IValidateFacesService interface {
	Validate(filePath string) (bool, error)
}
