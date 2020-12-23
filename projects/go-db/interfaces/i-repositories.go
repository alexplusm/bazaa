package interfaces

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
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
	InsertList(screenshots []dao.ScreenshotCreateDAO) error
	SelectListByGameID(gameID string) ([]dao.ScreenshotRetrieveDAO, error)
	Exist(screenshotID string) (bool, error)
	UpdateUsersAnswer(screenshotID, usersAnswer string) error
	CountByGame(gameID string) (int, error)
}

type IAnswerRepo interface {
	InsertOne(answer dao.AnswerInsertDAO) error
	InsertList(answers []dao.AnswerInsertDAO)
	SelectScreenshotResult(gameID, screenshotID string) ([]dao.AnswerScreenshotRetrieveDAO, error)
	SelectAnsweredScreenshotsByGame(gameID string) (dao.AnsweredScreenshotsDAO, error)
	SelectListTODO(gameID string, from, to time.Time) ([]dao.AnswerScreenshotRetrieveDAO, error)
	SelectListByUserAndGame(userID string, gameID string, from, to time.Time) ([]dao.AnswerScreenshotRetrieveDAO, error)
}

type IUserRepo interface {
	InsertOne(user dao.UserDAO) error
	Exist(userID string) (bool, error)
}
