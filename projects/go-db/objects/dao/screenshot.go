package dao

type ScreenshotDAO struct {
	Filename string
	GameID   string
	SourceID string
}

// TODO:refactor
type ScreenshotWithExpertAnswerDAO struct {
	ScreenshotDAO
	ExpertAnswer int
}
