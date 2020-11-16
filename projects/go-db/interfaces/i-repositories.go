package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type IGameRepository interface {
	InsertGame(game dao.GameDAO) (string, error)
	SelectGame(gameID string) (dao.GameDAO, error)
}

type IExtSystemRepository interface {
	InsertExtSystem(extSystemDAO dao.ExtSystemDAO) error
	SelectExtSystems() ([]dao.ExtSystemDAO, error)
}

type ISourceRepository interface {
	InsertSource(source dao.SourceDAO) (string, error)
}

type IScreenshotRepository interface {
	SelectScreenshotsByGameID(gameID string) ([]dao.ScreenshotDAOFull, error)
	InsertScreenshots(screenshots []dao.ScreenshotDAO) error
	InsertScreenshotsWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error
}

type IUserRepository interface {
	InsertUser(user dao.UserDAO) error
}
