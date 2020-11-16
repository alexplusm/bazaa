package interfaces

import (
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type IGameService interface {
	CreateGame(game bo.GameBO) (string, error)
	GetGame(gameID string) (bo.GameBO, error)
}

type IAttachSourceToGameService interface {
	AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error
	AttachSchedulesToGame(gameID string) error
}

type IExtSystemService interface {
	CreateExtSystem(extSystem bo.ExtSystemBO) error
}

type IGameCacheService interface {
	PrepareGame(gameID string) error
	GameWithSameExtSystemIDExist(gameID, extSystemID string) bool
}

type IScreenshotCacheService interface {
	CanSetUserAnswerToScreenshot(userID, screenshotID string) bool
	GetScreenshot(gameID, userID string) (dao.ScreenshotURLDAO, bool)
	GetUsersAnswers(screenshotID string) []bo.UserAnswerCacheBO
	SetUserAnswerToScreenshot(userID, screenshotID, answer string)
	ScreenshotExist(screenshotID string) bool
}

type IScreenshotUserAnswerService interface {
	BuildUserAnswerResponse(
		userID string, answersBO []bo.UserAnswerCacheBO,
	) dto.UserAnswerResponseData
	ScreenshotIsFinished(answers []bo.UserAnswerCacheBO) bool
}

type IUserService interface {
	CreateUser(userID string) error
}
