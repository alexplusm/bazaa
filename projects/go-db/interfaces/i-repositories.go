package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"time"
)

type IGameRepository interface {
	InsertGame(game dao.GameDAO) (string, error)
	SelectGame(gameID string) (dao.GameDAO, error)
	SelectGames(extSystemID string) ([]dao.GameDAO, error)
}

type IExtSystemRepository interface {
	InsertExtSystem(extSystemDAO dao.ExtSystemDAO) (string, error)
	SelectExtSystems() ([]dao.ExtSystemDAO, error)
	ExtSystemExist(extSystemID string) (bool, error)
}

type ISourceRepository interface {
	InsertSource(source dao.SourceDAO) (string, error)
}

type IScreenshotRepository interface {
	SelectScreenshotsByGameID(gameID string) ([]dao.ScreenshotDAOFull, error)
	InsertScreenshots(screenshots []dao.ScreenshotDAO) error
	InsertScreenshotsWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error
	UpdateScreenshotUsersAnswer(screenshotID, usersAnswer string) error
}

type IAnswerRepository interface {
	InsertAnswer(answer dao.AnswerDAO) error
	InsertAnswers(answers []dao.AnswerDAO)
	SelectAnswersByUser(
		userID string, gameIDs []string, from, to time.Time,
	) ([]dao.AnswerStatDAO, error)
}

type IUserRepository interface {
	InsertUser(user dao.UserDAO) error
	UserExist(userID string) (bool, error)
}
