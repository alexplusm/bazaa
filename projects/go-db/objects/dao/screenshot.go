package dao

type ScreenshotDAO struct {
	Filename string
	GameID   string
	SourceID string
}

type ScreenshotWithExpertAnswerDAO struct {
	ScreenshotDAO
	ExpertAnswer string
}

type ScreenshotDAOFull struct {
	ScreenshotID string
	GameID       string
	SourceID     string
	Filename     string
	ExpertAnswer string
	UsersAnswer  string
}
