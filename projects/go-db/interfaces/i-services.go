package interfaces

import (
	"mime/multipart"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type IGameService interface {
	CreateGame(game bo.GameBO) (string, error)
	GetGame(gameID string) (bo.GameBO, error)
	GetGames(extSystemID string) ([]bo.GameBO, error)
	GameExist(gameID string) (bool, error)
	FilterGames(gamesID []string, games []bo.GameBO) []bo.GameBO
	GetEarliestGame(games []bo.GameBO) bo.GameBO
}

type IAttachSourceToGameService interface {
	AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error
	AttachSchedulesToGame(gameID string) error
}

type IExtSystemService interface {
	CreateExtSystem(extSystem bo.ExtSystemBO) (string, error)
	ExtSystemExist(extSystemID string) (bool, error)
	ExtSystemList() ([]bo.ExtSystemBO, error)
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
	ABC(gameID string, from, to time.Time) ([]dao.AnswerStatLeadDAO, error)
}

type IScreenshotService interface {
	ScreenshotExist(screenshotID string) (bool, error)
	ScreenshotCountByGame(gameID string) (int, error)
}

type ISourceService interface {
	GetSourcesByGame(gameID string) ([]dao.Source2DAO, error)
}

type IActiveUsersService interface {
	SetUserActivity(gameID, userID string)
	CountOfActiveUsers(gameID string) (int, error)
}

type IDurationService interface {
	GetDurationByGame(from, to string, game bo.GameBO) (time.Time, time.Time)
}

type ILeaderboardService interface {
	GetLeaderboard(gameIDs []string, from, to time.Time, limit int) dto.LeadersResponseDTO
}
