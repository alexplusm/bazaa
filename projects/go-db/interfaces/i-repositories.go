package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"time"
)

type IGameRepository interface {
	InsertGame(game dao.GameDAO) (string, error)
	SelectGame(gameID string) (dao.GameDAO, error)
	SelectGames(extSystemID string) ([]dao.GameDAO, error)
	GameExist(gameID string) (bool, error)
}

type IExtSystemRepository interface {
	InsertExtSystem(extSystemDAO dao.ExtSystemDAO) (string, error)
	SelectExtSystems() ([]dao.ExtSystemDAO, error)
	ExtSystemExist(extSystemID string) (bool, error)
}

type ISourceRepository interface {
	InsertSource(source dao.SourceDAO) (string, error)
	SelectSourcesByGame(gameID string) ([]dao.Source2DAO, error)
}

type IScreenshotRepository interface {
	SelectScreenshotsByGameID(gameID string) ([]dao.ScreenshotDAOFull, error)
	InsertScreenshots(screenshots []dao.ScreenshotDAO) error
	InsertScreenshotsWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error
	UpdateScreenshotUsersAnswer(screenshotID, usersAnswer string) error
	ScreenshotExist(screenshotID string) (bool, error)
	ScreenshotCountByGame(gameID string) (int, error)
}

type IAnswerRepository interface {
	InsertAnswer(answer dao.AnswerDAO) error
	InsertAnswers(answers []dao.AnswerDAO)
	SelectScreenshotResult(gameID, screenshotID string) ([]dao.ScreenshotResultDAO, error)
	SelectAnsweredScreenshotsByGame(gameID string) (dao.AnsweredScreenshotsDAO, error)
	SelectAnswersTODO(gameID string, from, to time.Time) ([]dao.AnswerStatLeadDAO, error)
	SelectAnswersByUserAndGame(userID string, gameID string, from, to time.Time) ([]dao.UserAnswerDAO, error)
}

type IUserRepository interface {
	InsertUser(user dao.UserDAO) error
	UserExist(userID string) (bool, error)
}
