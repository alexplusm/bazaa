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
	InsertOne(source dao.SourceInsertDAO) (string, error)
	SelectListByGame(gameID string) ([]dao.SourceRetrieveDAO, error)
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
	InsertOne(answer dao.AnswerDAO) error
	InsertList(answers []dao.AnswerDAO)
	SelectScreenshotResult(gameID, screenshotID string) ([]dao.AnswerScreenshotRetrieveDAO, error)
	SelectAnsweredScreenshotsByGame(gameID string) (dao.AnsweredScreenshotsDAO, error)
	SelectListTODO(gameID string, from, to time.Time) ([]dao.AnswerScreenshotRetrieveDAO, error)
	SelectListByUserAndGame(userID string, gameID string, from, to time.Time) ([]dao.AnswerScreenshotRetrieveDAO, error)
}

type IUserRepo interface {
	InsertOne(user dao.UserDAO) error
	Exist(userID string) (bool, error)
}
