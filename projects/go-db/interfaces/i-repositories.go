package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"time"
)

type IGameRepo interface {
	InsertOne(game dao.GameDAO) (string, error)
	SelectOne(gameID string) (dao.GameDAO, error)
	SelectList(extSystemID string) ([]dao.GameDAO, error)
	Exist(gameID string) (bool, error)
}

type IExtSystemRepo interface {
	InsertOne(extSystemDAO dao.ExtSystemDAO) (string, error)
	SelectList() ([]dao.ExtSystemDAO, error)
	Exist(extSystemID string) (bool, error)
}

type ISourceRepo interface {
	InsertOne(source dao.SourceDAO) (string, error)
	SelectListByGame(gameID string) ([]dao.Source2DAO, error)
}

type IScreenshotRepo interface {
	SelectListByGameID(gameID string) ([]dao.ScreenshotDAOFull, error)
	InsertList(screenshots []dao.ScreenshotDAO) error
	InsertListWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error
	UpdateScreenshotUsersAnswer(screenshotID, usersAnswer string) error
	Exist(screenshotID string) (bool, error)
	ScreenshotCountByGame(gameID string) (int, error)
}

type IAnswerRepo interface {
	InsertAnswer(answer dao.AnswerDAO) error
	InsertAnswers(answers []dao.AnswerDAO)
	SelectScreenshotResult(gameID, screenshotID string) ([]dao.ScreenshotResultDAO, error)
	SelectAnsweredScreenshotsByGame(gameID string) (dao.AnsweredScreenshotsDAO, error)
	SelectAnswersTODO(gameID string, from, to time.Time) ([]dao.AnswerStatLeadDAO, error)
	SelectAnswersByUserAndGame(userID string, gameID string, from, to time.Time) ([]dao.UserAnswerDAO, error)
}

type IUserRepo interface {
	InsertOne(user dao.UserDAO) error
	Exist(userID string) (bool, error)
}
