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
